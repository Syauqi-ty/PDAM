createdb:
	docker exec -it pampam_postgres createdb -U postgres pampam

dropdb:
	docker exec -it pampam_postgres dropdb -U postgres pampam

run:
	docker start pampam_postgres

stop:
	docker stop pampam_postgres


.PHONY:createdb , dropdb, run, stop