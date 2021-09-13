DOCKER_COMPOSE_NAME_MIGRATION=goose

#Dump the migration status for the current DB
migr-status:
	docker-compose exec ${DOCKER_COMPOSE_NAME_MIGRATION} sh -c "./goose -dir=./migrations/ -v status"

#Migrate the DB to the most recent version available
migr-up:
	docker-compose exec ${DOCKER_COMPOSE_NAME_MIGRATION} sh -c "./goose -dir=./migrations/ -v up"

#Roll back the version by 1
migr-down:
	docker-compose exec ${DOCKER_COMPOSE_NAME_MIGRATION} sh -c "./goose -dir=./migrations/ -v down"

#Re-run the latest migration
migr-redo:
	docker-compose exec ${DOCKER_COMPOSE_NAME_MIGRATION} sh -c "./goose -dir=./migrations/ -v redo"

#Creates new migration file with the current timestamp
migr-create:
	docker-compose exec ${DOCKER_COMPOSE_NAME_MIGRATION} sh -c "./goose -dir=./migrations/ -v create \"\" sql"
