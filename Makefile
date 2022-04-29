rebuild:
	@docker compose down
	@docker compose build --no-cache
	@docker compose up -d

reload-app:
	@cd ./app; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags="all=-N -l" -o ./main;
	@$(eval pid := $(shell docker ps -qf "Name=automatic-trade_app"))
	@docker cp ./app/main $(pid):/app/main
	@docker restart $(pid)
	@rm -f ./app/main

exec:
	@docker compose exec app go run .