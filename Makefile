install:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin
	go mod download

run:
	go run ./Service/cmd/gometr/main.go

lint:
	golangci-lint run ./... -v

docker-build:
	docker build -f ./Service/build/gometr/Dockerfile -t gometr:v1 . 

docker-run:
	docker run -d --name gometr -p 8001:8000 gometr:v1

task684:
	docker-compose -f Exercises/5-6.Go_in_practice/3/deployments/task1/docker-compose.yml up -d --build

task684wrk:
	for rps in 100 500 1000 1500 2000 ; do \
	    echo "RPS=$$rps" ; \
		docker run --rm -it -v $(shell pwd)/Exercises/5-6.Go_in_practice/3/deployments/task1/scripts:/scripts --network=task1_messagetask1 artfrela/wrk2:alpine3.19 -t4 -c16 -d15s -R$$rps -s /scripts/rate.lua http://proxy:8080 ; \
	done

clean684:
	$(eval IMAGES := $(shell docker-compose -f Exercises/5-6.Go_in_practice/3/deployments/task1/docker-compose.yml images | awk '/task1/ { print $$2" "$$4 }' | grep 'task1-' | awk '{ print $$2 }'))
	docker-compose -f Exercises/5-6.Go_in_practice/3/deployments/task1/docker-compose.yml rm --stop --force -v
	docker rmi $(IMAGES) --force 2>/dev/null; true

task684nats:
	docker-compose -f Exercises/5-6.Go_in_practice/3/deployments/task1/docker-compose_nats.yml up -d --build

clean684nats:
	$(eval IMAGES := $(shell docker-compose -f Exercises/5-6.Go_in_practice/3/deployments/task1/docker-compose_nats.yml images | awk '/task1/ { print $$2" "$$4 }' | grep 'task1-' | awk '{ print $$2 }'))
	docker-compose -f Exercises/5-6.Go_in_practice/3/deployments/task1/docker-compose_nats.yml rm --stop --force -v
	docker rmi $(IMAGES) --force 2>/dev/null; true