run:
	@docker compose up -d
	@docker compose -f docker-compose.yml --profile tools run --rm migrate up 
stop:
	@docker compose down
.PHONY: run stop
