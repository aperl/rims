

.PHONY: build
build:
	docker build --target release -t allenperl/rims .
	docker tag allenperl/rims:latest allenperl/rims:0.0.3
