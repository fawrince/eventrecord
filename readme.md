# What is this about
1. Client application records the mouse movement from the screen
2. Sends coordinate to the server over WebSocket connection
3. Server emits a kafka message with coordinates to the topic specified on container startup
4. When replay recording requested (via button) - server sends coordinates to the client over sse-connection
5. Client replays mouse movement history on the screen

# How to run
1. docker build --tag eventrecord-server:1.0.0 -f Dockerfile-server .
2. docker-compose up

localhost:1001 - Web Server\
localhost:3000 - Grafana \
localhost:9090 - Prometheus \
localhost:5601 - Kibana 

# Components used
github.com/Shopify/sarama - go client for Kafka\
github.com/gofiber/fiber/v2 - fast webserver based on fasthttp\
github.com/prometheus/client_golang - go client for prometheus\
github.com/sirupsen/logrus - well-known logger\
google.golang.org/grpc - grpc\
google.golang.org/protobuf - protobuf

# Hints
### build docker image
docker build --tag eventrecord-server:1.0.0 -f Dockerfile-server .

### enter docker daemon
screen ~/Library/Containers/com.docker.docker/Data/vms/0/tty

### push image to gcloud
docker push europe-north1-docker.pkg.dev/reference-yen-341621/dockers/eventrecord-server:1.0.0
docker push europe-north1-docker.pkg.dev/reference-yen-341621/dockers/wurstmeister/kafka:latest

### assign tag to a container in gcloud
docker tag eventrecord-server:1.0.0 europe-north1-docker.pkg.dev/reference-yen-341621/dockers/eventrecord-server:1.0.0
docker tag wurstmeister/kafka:latest europe-north1-docker.pkg.dev/reference-yen-341621/dockers/wurstmeister/kafka:latest

### run kafka console consumer
docker exec -it broker  kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic coordinates  --from-beginning --partition=

### apply kubectl deployment
kubectl apply -f zookeeper-service.yaml,zookeeper-deployment.yaml,kafka-service.yaml,\
kafka-deployment.yaml,kafka-claim0-persistentvolumeclaim.yaml,\
eventrecord-server-service.yaml,eventrecord-server-deployment.yaml

### run the protoc generation
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
grpc/coordinate_transporter.proto
