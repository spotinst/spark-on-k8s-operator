REGISTRY ?=
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)

VERSION = $(shell git rev-parse --short=7 HEAD)
IMAGE_NAME = spark-operator
SPARK_BASE_IMAGE = gcr.io/ocean-spark/spark:platform-3.5-gen25
TAG = $(VERSION)-gen24
SPARK_BASE_IMAGE_RH = public.ecr.aws/f4k1p1n4/spark:netapp-spark-support-image-v3.5.3-rel20241122
TAG_RH = $(VERSION)-ubi9-rel20241122

build_linux:
	docker buildx build --load \
		-t $(REGISTRY)$(IMAGE_NAME):$(TAG) \
		--build-arg SPARK_IMAGE=$(SPARK_BASE_IMAGE) .

build_rh:
	docker buildx build --load \
		-t $(REGISTRY)$(IMAGE_NAME):$(TAG_RH) \
		-f Dockerfile.rh \
		--build-arg SPARK_IMAGE=$(SPARK_BASE_IMAGE_RH) .

build: build_linux build_rh

push_linux:
	docker tag $(REGISTRY)$(IMAGE_NAME):$(TAG) $(REGISTRY)$(IMAGE_NAME):${BRANCH}
	docker push $(REGISTRY)$(IMAGE_NAME):${BRANCH}
	@if [ "$(BRANCH)" = "ocean-spark" ]; then \
		docker push $(REGISTRY)$(IMAGE_NAME):$(TAG); \
	fi

push_rh:
	docker tag $(REGISTRY)$(IMAGE_NAME):$(TAG_RH) $(REGISTRY)$(IMAGE_NAME):${BRANCH}-rh
	docker push $(REGISTRY)$(IMAGE_NAME):${BRANCH}-rh
	@if [ "$(BRANCH)" = "ocean-spark" ]; then \
		docker push $(REGISTRY)$(IMAGE_NAME):$(TAG_RH); \
	fi

push: push_linux push_rh

show:
	@echo "Images available:"
	@if [ "$(BRANCH)" = "ocean-spark" ]; then \
		echo "-> $(REGISTRY)$(IMAGE_NAME):$(TAG)"; \
		echo "-> $(REGISTRY)$(IMAGE_NAME):$(TAG_RH)"; \
	fi
	@echo "-> $(REGISTRY)$(IMAGE_NAME):${BRANCH}"
	@echo "-> $(REGISTRY)$(IMAGE_NAME):${BRANCH}-rh"
