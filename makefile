mig-run:
	migrate create -ext sql -dir migrations -seq create_table

mig-up:
	migrate -database 'postgres://sayyidmuhammad:root@localhost:5432/shopp?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://sayyidmuhammad:root@localhost:5432/shopp?sslmode=disable' -path migrations down
