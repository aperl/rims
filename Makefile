VERSION=0.1.0
CONTAINER_NAME=allenperl/rims

.PHONY: build
build:
	docker build --target release -t $(CONTAINER_NAME) .
	docker tag $(CONTAINER_NAME):latest $(CONTAINER_NAME):$(VERSION)

publish:
	docker push $(CONTAINER_NAME):latest
	docker push $(CONTAINER_NAME):$(VERSION)
