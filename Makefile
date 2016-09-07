install:
	docker build --tag=marmelab-go .

run:
	@docker run --rm --volume="`pwd`:/srv" -ti marmelab-go run src/main/main.go
