.PHONY: docker-run

docker-build:
	docker buildx build -t basse .

docker-run:
	docker run --rm -v $$(pwd):/base -v /base/tmp -p 8003:8002 --name basse-air basse