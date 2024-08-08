module github.com/GoogleCloudPlatform/spark-on-k8s-operator

go 1.22.5

require (
	cloud.google.com/go/storage v1.30.1
	github.com/aws/aws-sdk-go v1.38.49
	github.com/evanphx/json-patch v4.12.0+incompatible
	github.com/golang/glog v1.2.1
	github.com/google/go-cloud v0.1.1
	github.com/google/uuid v1.6.0
	github.com/olekukonko/tablewriter v0.0.4
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.16.0
	github.com/prometheus/client_model v0.4.0
	github.com/robfig/cron v1.2.0
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/net v0.28.0
	golang.org/x/sync v0.8.0
	golang.org/x/time v0.6.0
	k8s.io/api v0.30.3
	k8s.io/apiextensions-apiserver v0.30.3
	k8s.io/apimachinery v0.30.3
	k8s.io/client-go v1.5.2
	k8s.io/kubectl v0.30.3
	k8s.io/kubernetes v1.30.3
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8
	volcano.sh/apis v1.1.0
)

require (
	cloud.google.com/go v0.110.6 // indirect
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	cloud.google.com/go/iam v1.1.1 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.11.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-runewidth v0.0.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/term v0.0.0-20221205130635-1aeaba878587 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xlab/treeprint v1.2.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.starlark.net v0.0.0-20230525235612-a134d8f9ddca // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/oauth2 v0.20.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/term v0.23.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.126.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/cli-runtime v0.30.3 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.13.5-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/kustomize/kyaml v0.14.3-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.29.3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.29.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.29.3
	k8s.io/apiserver => k8s.io/apiserver v0.29.3
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.29.3
	k8s.io/client-go => k8s.io/client-go v0.29.3
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.29.3
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.29.3
	k8s.io/code-generator => k8s.io/code-generator v0.29.3
	k8s.io/component-base => k8s.io/component-base v0.29.3
	k8s.io/cri-api => k8s.io/cri-api v0.29.3
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.29.3
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.29.3
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.29.3
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.29.3
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.29.3
	k8s.io/kubectl => k8s.io/kubectl v0.29.3
	k8s.io/kubelet => k8s.io/kubelet v0.29.3
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.29.3
	k8s.io/metrics => k8s.io/metrics v0.29.3
	k8s.io/node-api => k8s.io/node-api v0.29.3
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.29.3
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.29.3
	k8s.io/sample-controller => k8s.io/sample-controller v0.29.3
)
