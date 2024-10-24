### Run server
compose-up:
	docker-compose up --build -d && docker-compose logs -f
.PHONY: compose-up

### Stop server
compose-down:
	docker-compose down --remove-orphans
.PHONY: compose-down