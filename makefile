RUN = docker compose exec api

up:
	docker compose up -d

migration/new:
	${RUN} sql-migrate new ${FILE_NAME}

migration/up:
	${RUN} sql-migrate up --env="local"

migration/down:
	${RUN} sql-migrate down --env="local"