package sparkapplication

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	kubeclientfake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"

	"github.com/kubeflow/spark-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	crdclientfake "github.com/kubeflow/spark-operator/pkg/client/clientset/versioned/fake"
	crdinformers "github.com/kubeflow/spark-operator/pkg/client/informers/externalversions"
	"github.com/kubeflow/spark-operator/pkg/util"
)

// separate this test with the original `controller_test.go` to simplify rebasing process in the future

// newAnotherFakeController is a copy of the function from the original controller_test
// except we don't enable the UIService (Behavior from ofas spark-operator)
func newAnotherFakeController(app *v1beta2.SparkApplication, pods ...*apiv1.Pod) (*Controller, *record.FakeRecorder) {
	crdclientfake.AddToScheme(scheme.Scheme)
	crdClient := crdclientfake.NewSimpleClientset()
	kubeClient := kubeclientfake.NewSimpleClientset()
	util.IngressCapabilities = map[string]bool{"networking.k8s.io/v1": true}
	informerFactory := crdinformers.NewSharedInformerFactory(crdClient, 0*time.Second)
	recorder := record.NewFakeRecorder(3)

	kubeClient.CoreV1().Nodes().Create(context.TODO(), &apiv1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node1",
		},
		Status: apiv1.NodeStatus{
			Addresses: []apiv1.NodeAddress{
				{
					Type:    apiv1.NodeExternalIP,
					Address: "12.34.56.78",
				},
			},
		},
	}, metav1.CreateOptions{})

	podInformerFactory := informers.NewSharedInformerFactory(kubeClient, 0*time.Second)
	controller := newSparkApplicationController(crdClient, kubeClient, informerFactory, podInformerFactory, recorder,
		&util.MetricConfig{}, "", "", nil, false, false, util.RatelimitConfig{}, 5)

	informer := informerFactory.Sparkoperator().V1beta2().SparkApplications().Informer()
	if app != nil {
		informer.GetIndexer().Add(app)
	}

	podInformer := podInformerFactory.Core().V1().Pods().Informer()
	for _, pod := range pods {
		if pod != nil {
			podInformer.GetIndexer().Add(pod)
		}
	}
	return controller, recorder
}

func TestSyncSparkApplication_When_Submission_Successes(t *testing.T) {

	/*
		test normal case when the submission is successes
		we received 2 times the NewState state for each CRD
		- first time, the controller will submit the application
		- second time, the controller should skip the submission
		Check submission is done only once and the expected state is SubmittedState
	*/

	originalSparkHome := os.Getenv(sparkHomeEnvVar)
	originalKubernetesServiceHost := os.Getenv(kubernetesServiceHostEnvVar)
	originalKubernetesServicePort := os.Getenv(kubernetesServicePortEnvVar)
	os.Setenv(sparkHomeEnvVar, "/spark")
	os.Setenv(kubernetesServiceHostEnvVar, "localhost")
	os.Setenv(kubernetesServicePortEnvVar, "443")
	defer func() {
		os.Setenv(sparkHomeEnvVar, originalSparkHome)
		os.Setenv(kubernetesServiceHostEnvVar, originalKubernetesServiceHost)
		os.Setenv(kubernetesServicePortEnvVar, originalKubernetesServicePort)
	}()

	restartPolicyNever := v1beta2.RestartPolicy{
		Type: v1beta2.Never,
	}

	// Create a new SparkApplication with NewState
	app := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: v1beta2.SparkApplicationSpec{
			RestartPolicy: restartPolicyNever,
		},
		Status: v1beta2.SparkApplicationStatus{
			AppState: v1beta2.ApplicationState{
				State: v1beta2.NewState,
			},
		},
	}

	ctrl, _ := newAnotherFakeController(app)
	_, err := ctrl.crdClient.SparkoperatorV1beta2().SparkApplications(app.Namespace).Create(context.TODO(), app, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}

	// Mock the execCommand to return a success
	execCommand = func(command string, args ...string) *exec.Cmd {
		cs := []string{"-test.run=TestHelperProcessSuccess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
		return cmd
	}

	// simulate the first NewState
	err = ctrl.syncSparkApplication(fmt.Sprintf("%s/%s", app.Namespace, app.Name))
	assert.Nil(t, err)
	updatedApp, err := ctrl.crdClient.SparkoperatorV1beta2().SparkApplications(app.Namespace).Get(context.TODO(), app.Name, metav1.GetOptions{})
	assert.Nil(t, err)
	assert.Equal(t, v1beta2.SubmittedState, updatedApp.Status.AppState.State)
	assert.Equal(t, float64(1), fetchCounterValue(ctrl.metrics.sparkAppSubmitCount, map[string]string{}))

	// simulate the second NewState (should skip the submission)
	err = ctrl.syncSparkApplication(fmt.Sprintf("%s/%s", app.Namespace, app.Name))
	assert.Nil(t, err)
	updatedApp, err = ctrl.crdClient.SparkoperatorV1beta2().SparkApplications(app.Namespace).Get(context.TODO(), app.Name, metav1.GetOptions{})
	assert.Nil(t, err)
	// check the state is still submitted
	assert.Equal(t, v1beta2.SubmittedState, updatedApp.Status.AppState.State)
	// check the submit count does not change
	assert.Equal(t, float64(1), fetchCounterValue(ctrl.metrics.sparkAppSubmitCount, map[string]string{}))
}

func TestSyncSparkApplication_When_Submission_Fails(t *testing.T) {
	/*
		test case when the submission is failed the first time
		we received 2 times the NewState state for each CRD
		- first time, the controller will submit the application and failed
		- second time, the controller should skip the submission
		Check submission is done only once and the expected state is FailedSubmissionState
	*/
	originalSparkHome := os.Getenv(sparkHomeEnvVar)
	originalKubernetesServiceHost := os.Getenv(kubernetesServiceHostEnvVar)
	originalKubernetesServicePort := os.Getenv(kubernetesServicePortEnvVar)
	os.Setenv(sparkHomeEnvVar, "/spark")
	os.Setenv(kubernetesServiceHostEnvVar, "localhost")
	os.Setenv(kubernetesServicePortEnvVar, "443")
	defer func() {
		os.Setenv(sparkHomeEnvVar, originalSparkHome)
		os.Setenv(kubernetesServiceHostEnvVar, originalKubernetesServiceHost)
		os.Setenv(kubernetesServicePortEnvVar, originalKubernetesServicePort)
	}()

	restartPolicyNever := v1beta2.RestartPolicy{
		Type: v1beta2.Never,
	}

	// Create a new SparkApplication with NewState
	app := &v1beta2.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: v1beta2.SparkApplicationSpec{
			RestartPolicy: restartPolicyNever,
		},
		Status: v1beta2.SparkApplicationStatus{
			AppState: v1beta2.ApplicationState{
				State: v1beta2.NewState,
			},
		},
	}

	ctrl, _ := newAnotherFakeController(app)
	_, err := ctrl.crdClient.SparkoperatorV1beta2().SparkApplications(app.Namespace).Create(context.TODO(), app, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}

	// Mock the execCommand to return a failure
	execCommand = func(command string, args ...string) *exec.Cmd {
		cmd := exec.Command("/bin/should-fail")
		return cmd
	}

	err = ctrl.syncSparkApplication(fmt.Sprintf("%s/%s", app.Namespace, app.Name))
	assert.Nil(t, err)
	updatedApp, err := ctrl.crdClient.SparkoperatorV1beta2().SparkApplications(app.Namespace).Get(context.TODO(), app.Name, metav1.GetOptions{})
	assert.Nil(t, err)
	assert.Equal(t, v1beta2.FailedSubmissionState, updatedApp.Status.AppState.State)
	assert.Equal(t, float64(0), fetchCounterValue(ctrl.metrics.sparkAppSubmitCount, map[string]string{}))
	assert.Equal(t, float64(1), fetchCounterValue(ctrl.metrics.sparkAppFailedSubmissionCount, map[string]string{}))

	// simulate the second NewState (should skip the submission)

	// This time, mock the command to be successful, but we expected the command is not executed
	execCommand = func(command string, args ...string) *exec.Cmd {
		cs := []string{"-test.run=TestHelperProcessSuccess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
		return cmd
	}
	err = ctrl.syncSparkApplication(fmt.Sprintf("%s/%s", app.Namespace, app.Name))
	assert.Nil(t, err)
	updatedApp, err = ctrl.crdClient.SparkoperatorV1beta2().SparkApplications(app.Namespace).Get(context.TODO(), app.Name, metav1.GetOptions{})
	assert.Nil(t, err)
	// check the CR state is still failedSubmission
	assert.Equal(t, v1beta2.FailedSubmissionState, updatedApp.Status.AppState.State)
	// check the submit count does not change
	assert.Equal(t, float64(0), fetchCounterValue(ctrl.metrics.sparkAppSubmitCount, map[string]string{}))
	assert.Equal(t, float64(1), fetchCounterValue(ctrl.metrics.sparkAppFailedSubmissionCount, map[string]string{}))
}
