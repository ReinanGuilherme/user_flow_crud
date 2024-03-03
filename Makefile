# Variáveis
DB_USER := root
DB_PASSWORD := secret
DB_NAME := user_flow
DB_HOST := localhost
DB_PORT := 5432
DB_SSL_MODE := disable
MIGRATIONS_PATH := migrations

# Comando para criar o banco de dados
create_db:
	docker exec -it postgres12 createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

# Comando para deletar o banco de dados
drop_db:
	docker exec -it postgres12 dropdb $(DB_NAME)

# Comando para inicializar as migrações
migrate_init:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq init_schema

# Comando para executar as migrações para cima
migrate_up:
	migrate -path $(MIGRATIONS_PATH) -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" up

# Comando para reverter as migrações para baixo
migrate_down:
	migrate -path $(MIGRATIONS_PATH) -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" down

# Comando para gerar as consultas golang a partir de uma query sql
sqlc:
	sqlc generate

.PHONY: migration_init create_db drop_db migrate_up migrate_down sqlc
