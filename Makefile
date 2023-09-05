migrate :
	go build ./cmd/migrator.go
db-up:
	go run ./cmd/migrator.go up
db-redo:
	go run cmd/migrator.go redo
db-create:
	go ./cmd/migrator.go create
db-status:
	go run ./cmd/migrator.go status
