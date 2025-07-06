.PHONY: docker-run docker-build

docker-build:
	docker buildx build -t basse .

docker-run:
	docker run --rm -v $$(pwd):/base -v /base/tmp -p 8002:8002 --name basse-air basse