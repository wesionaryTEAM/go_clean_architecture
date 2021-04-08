include .env
MIGRATE=docker-compose exec web migrate -path=migration -database "mysql://${DBUsername}:${DBPassword}@tcp(${DBHost}:${DBPort})/${DBName}" -verbose

check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

setup:
	git config core.hooksPath hooks

dev:
		gin appPort 5000 -i run server.go
migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -dir migration  $$NAME

.PHONY: migrate-up migrate-down force goto drop create

tags:
	@read -r -p "Which file? " FILE; \
	read -r -p "Which struct? " STRUCT; \
	read -r -p "Which tag? " TAG; \
	gomodifytags -file $$FILE -struct $$STRUCT -add-tags $$TAG -w
