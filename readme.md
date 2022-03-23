# what is this about
1. Application records the mouse movement from the screen
2. Sends it to the web server over http
3. Emits a kafka message with mouse coordinates on the backend
4. If reqeusted (via button) - sends coordinates to the client over sse-connection
5. Replays mouse movement history on the screen

# build
docker build --tag eventrecord-server:1.0.0 -f Dockerfile-server .
docker build --tag eventrecord-consumer:1.0.0 -f Dockerfile-consumer .

# build for a scratch image
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# enter docker daemon
screen ~/Library/Containers/com.docker.docker/Data/vms/0/tty

# push image to gcloud
docker push europe-north1-docker.pkg.dev/reference-yen-341621/dockers/invest:1.0.0

# assign tag to a container in gcloud
docker tag invest:1.0.0 europe-north1-docker.pkg.dev/reference-yen-341621/dockers/invest:1.0.0

# run kafka console consumer
docker exec -it broker  kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic coordinates  --from-beginning --partition=


