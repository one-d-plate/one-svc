create:
	mkdir -p migrations
	goose -dir migrations create $(name) sql

migrate:
	goose -dir migrations mysql "user:secret@tcp(localhost:3306)/auth" up
	
rollback:
	goose -dir migrations mysql "user:secret@tcp(127.0.0.1:3306)/auth" down