migrate_up:
	migrate -database "postgresql://postgres:postgres123@localhost:5432/guesthouse-rental?sslmode=disable" -path schema/migrations up


.PHONY:
	migrate_up