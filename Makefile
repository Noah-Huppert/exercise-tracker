.PHONY: run mongo

# run main.go
run:
	go run main.go

# mongo runs a local Mongo server in a Docker container
mongo:
	mkdir -p container-data/mongo
	docker run \
		-it \
		--rm \
		--name exercise-tracker-mongo \
		--net host \
		-v "${PWD}"/container-data/mongo:/data/db \
		mongo
