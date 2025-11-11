include .env
export

MIGRATE = atlas migrate

DB_URL := postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_FORWARD_PORT)/$(DB_NAME)?sslmode=disable
DEV_URL := docker://postgres/17/dev?search_path=public

migrate-status:
	$(MIGRATE) status --url "${DB_URL}" --dir "file://migrations"

migrate-diff:
	$(MIGRATE) diff --env gorm

migrate-apply:
	$(MIGRATE) apply --url "${DB_URL}" --allow-dirty --exec-order non-linear

migrate-down:
	$(MIGRATE) down --url "${DB_URL}" --dev-url "${DEV_URL}"

migrate-hash:
	$(MIGRATE) hash

lint-setup:
	python3 -m ensurepip --upgrade
	sudo pip3 install pre-commit
	pre-commit install
	pre-commit autoupdate

.PHONY: migrate-status migrate-diff migrate-apply migrate-down migrate-hash lint-setup