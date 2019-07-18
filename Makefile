.PHONY: build_run
build_run: 
	docker build -t jdr .
	docker run -d --name jdr -p 9300:9300 --rm --env PORT=9300 jdr

.PHONY: stop
stop:
	docker stop jdr

.PHONY: rebuild
rebuild: 
	make stop
	make build_run
