source: https://ewanvalentine.io/microservices-in-golang-part-2/

### Keywords
docker, go-micro

### Build docker image
sudo docker build -t shippy-service-consignment .
sudo docker build -t shippy-cli-consignment .

### Create a container from server image
sudo docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment
sudo docker run shippy-cli-consignment

### Update protobuf 
protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto