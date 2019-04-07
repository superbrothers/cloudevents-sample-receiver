.PHONY: build
build:
		CGO_ENABLED=0 go build -a -installsuffix cgo -o out/cloudevents-sample-receiver main.go

.PHONY: image
image:
		./hack/build.sh
