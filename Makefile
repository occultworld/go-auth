s:
	docker compose run --rm --use-aliases --service-ports app air
	${MAKE} clean-on-exit

shell:
	docker compose run --rm --use-aliases app bash
	${MAKE} clean-on-exit

db-shell:
	docker compose run --rm --use-aliases -e PGPASSWORD=password db psql -Upostgres -hdb golang_auth