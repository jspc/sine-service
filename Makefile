include .ori-env/helm-charts/Makefile

CIRCLE_SHA1 ?= local
APP = sine-service

default: all

all: service/sine-service.pb.go sine-service

clean:
	-rm service/*.pb.go sine-service

service/sine-service.pb.go:
	protoc -I grpc/ grpc/sine-service.proto --go_out=plugins=grpc:service

sine-service: service/sine-service.pb.go
	CGO_ENABLED=0 go build -o sine-service

.PHONY: docker
docker: sine-service
	docker build -t jspc/sine-service:latest -t jspc/sine-service:${CIRCLE_SHA1} .
	docker push jspc/sine-service
