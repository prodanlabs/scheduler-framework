TAG?=latest

all: build

build:
        CGO_ENABLED=0 GOOS=linux go build -o app

image:
        docker build  -t prodan/scheduler-framework:$(TAG).

push:
        docker push prodan/scheduler-framework:$(TAG)