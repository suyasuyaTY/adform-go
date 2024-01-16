CMD := docker-compose

up:
	@$(CMD) up -d

run:
	@$(CMD) exec app go run main.go

down:
	@$(CMD) down

fmt:
	@$(CMD) exec app go fmt ./...

clean:
	@$(CMD) down --rmi all --volumes --remove-orphans