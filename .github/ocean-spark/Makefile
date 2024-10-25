REGISTRY ?=
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)

VERSION = $(shell git rev-parse --short=7 HEAD)
IMAGE_NAME = spark-operator
SPARK_BASE_IMAGE = public.ecr.aws/f4k1p1n4/spark:netapp-spark-support-image-v3.5-rel20240916
TAG = $(VERSION)
PLATFORM ?= amd64


build:
	docker buildx build --load \
		-t $(REGISTRY)$(IMAGE_NAME):$(TAG)-$(PLATFORM) \
		-f .github/ocean-spark/Dockerfile \
		--platform linux/$(PLATFORM) \
		--build-arg SPARK_IMAGE=$(SPARK_BASE_IMAGE) .

push:
	docker tag $(REGISTRY)$(IMAGE_NAME):$(TAG)-$(PLATFORM) $(REGISTRY)$(IMAGE_NAME):${BRANCH}-$(PLATFORM)
	docker push $(REGISTRY)$(IMAGE_NAME):${BRANCH}-$(PLATFORM)
	@if [ "$(BRANCH)" = "ocean-spark-v2" ]; then \
		docker push $(REGISTRY)$(IMAGE_NAME):$(TAG)-$(PLATFORM); \
	fi


push-manifest:
	docker manifest create $(REGISTRY)$(IMAGE_NAME):${BRANCH} \
		$(REGISTRY)$(IMAGE_NAME):${BRANCH}-amd64 \
		$(REGISTRY)$(IMAGE_NAME):${BRANCH}-arm64
	docker manifest push $(REGISTRY)$(IMAGE_NAME):${BRANCH}
	@if [ "$(BRANCH)" = "ocean-spark-v2" ]; then \
		docker manifest create $(REGISTRY)$(IMAGE_NAME):$(TAG) \
			$(REGISTRY)$(IMAGE_NAME):$(TAG)-amd64 \
			$(REGISTRY)$(IMAGE_NAME):$(TAG)-arm64; \
		docker manifest push $(REGISTRY)$(IMAGE_NAME):$(TAG); \
	fi

show:
	@echo "Images available:"
	@if [ "$(BRANCH)" = "ocean-spark-v2" ]; then \
		echo "-> $(REGISTRY)$(IMAGE_NAME):$(TAG)"; \
	fi
	@echo "-> $(REGISTRY)$(IMAGE_NAME):${BRANCH}"