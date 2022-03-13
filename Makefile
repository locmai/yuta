messaging:
	go run components/messaging/main.go --config ./components/messaging/config.yaml

core:
	go run components/core/main.go --config ./components/core/config.yaml

build: build-messaging build-core

.PHONY: build-messaging
build-messaging:
	CGO_ENABLED=0 GOOS=linux go build -o ./build ./components/messaging 
	docker build -t locmai/yuta-messaging --build-arg component=messaging .

.PHONY: build-core
build-core:
	CGO_ENABLED=0 GOOS=linux go build -o ./build ./components/core 
	docker build -t locmai/yuta-core --build-arg component=core .
