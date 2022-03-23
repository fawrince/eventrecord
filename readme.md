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
docker push europe-north1-docker.pkg.dev/reference-yen-341621/dockers/invest:1.0.0

### assign tag to a container in gcloud
docker tag invest:1.0.0 europe-north1-docker.pkg.dev/reference-yen-341621/dockers/invest:1.0.0

### run kafka console consumer
docker exec -it broker  kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic coordinates  --from-beginning --partition=


