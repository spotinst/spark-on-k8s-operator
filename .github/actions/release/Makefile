REGISTRY ?=
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)

VERSION = $(shell git rev-parse --short=7 HEAD)
IMAGE_NAME = spark-operator
SPARK_BASE_IMAGE = gcr.io/datamechanics/spark:platform-3.2.1-dm18
TAG = $(VERSION)-dm18

build:
	docker buildx build --load \
		-t $(REGISTRY)$(IMAGE_NAME):$(TAG) \
		--build-arg SPARK_IMAGE=$(SPARK_BASE_IMAGE) .

push:
	docker tag $(REGISTRY)$(IMAGE_NAME):$(TAG) $(REGISTRY)$(IMAGE_NAME):${BRANCH}
	docker push $(REGISTRY)$(IMAGE_NAME):${BRANCH}
	@if [ "$(BRANCH)" = "ocean-spark" ]; then \
		docker push $(REGISTRY)$(IMAGE_NAME):$(TAG); \
	fi

show:
	@echo "Images available:"
	@if [ "$(BRANCH)" = "ocean-spark" ]; then \
		echo "-> $(REGISTRY)$(IMAGE_NAME):$(TAG)"; \
	fi
	@echo "-> $(REGISTRY)$(IMAGE_NAME):${BRANCH}"
