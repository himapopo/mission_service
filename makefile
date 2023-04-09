RUN = docker compose exec api

up:
	docker compose up -d

stop:
	docker compose stop

migration/new:
	${RUN} sql-migrate new ${FILE_NAME}

migration/up:
	${RUN} sql-migrate up --env="local"

migration/down:
	${RUN} sql-migrate down --env="local"

sqlboiler:
	${RUN} sqlboiler psql

test-data-init:
	docker compose cp ./testdata/test.sql db:/
	docker compose exec db psql -h localhost -p 5432 -U root -d mission_service -f test.sql