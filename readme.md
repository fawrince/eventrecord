# What is this about
1. Client application records the mouse movement from the screen
2. Send coordinates to the web server over http
3. Server-side emits a kafka message with coordinates to the topic specified on container startup
4. When replay recording requested (via button) - sends coordinates to the client over sse-connection
5. Replays mouse movement history on the screen

# Prompts
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

protoc --js_out=library=grpc/coordinate_transporter,binary:static/js \
grpc/coordinate_transporter.proto

protoc \
--grpc-web_out=import_style=commonjs,mode=grpcweb:static/js \
--js_out=import_style=commonjs,binary:static/js \
grpc/coordinate_transporter.proto

### run browserify to 
npx browserify static/js/grpc/coordinate_transporter_pb.js -o ct_pb.js
npx browserify static/js/grpc/coordinate_transporter_grpc_web_pb.js -o ct_web_pb.js

