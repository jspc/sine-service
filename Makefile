default: all

all: service/sine-service.pb.go

clean:
	-rm service/*.pb.go

service/sine-service.pb.go:
	protoc -I grpc/ grpc/sine-service.proto --go_out=plugins=grpc:service