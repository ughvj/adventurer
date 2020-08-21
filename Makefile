IMAGE=adventurer
CONTAINER=container_adventurer

build:
	docker build -t $(IMAGE) .

run:
	docker run \
		-v `pwd`/src:/go/src/adventurer \
		$(IMAGE) go run main.go

shell:
	docker run -it \
		-v `pwd`/src:/go/src/adventurer \
		$(IMAGE) /bin/bash

rm:
	docker rm -f $(CONTAINER)

clean:rm
	docker rmi $(IMAGE)