IMAGE=adventurer
CONTAINER=container_adventurer

build:
	docker build -t $(IMAGE) .

run:
	docker run \
		-v `pwd`/src:/go/src/github.com/ughvj/adventurer \
		--env-file ./vars.env \
		$(IMAGE) go run main.go

rm:
	docker rm -f $(CONTAINER)

clean:rm
	docker rmi $(IMAGE)