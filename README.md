source: https://ewanvalentine.io/microservices-in-golang-part-1/

### Keywords
grpc, protobuf

### To generate protobuf code:
protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto