clean:
	rm edge-test-tcp-udp-server

build:
	go build

build_docker_image:
	docker build -t edge-test-tcp-udp-server .

push_docker_image:
	docker tag edge-test-tcp-udp-server image-hub.simplifyd.dev/edge/edge-test-tcp-udp-server
	docker push image-hub.simplifyd.dev/edge/edge-test-tcp-udp-server

build_and_push_docker_image:
	docker build -t edge-test-tcp-udp-server .
	docker tag edge-test-tcp-udp-server image-hub.simplifyd.dev/edge/edge-test-tcp-udp-server
	docker push image-hub.simplifyd.dev/edge/edge-test-tcp-udp-server

linux:
	GOOS=linux go build

win:
	GOOS=windows go build -o ./bin

.PHONY: proto_gen clean build