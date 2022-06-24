run:
	PORT=3000 \
	POSTGRES_CONNECTION_URL="postgres://cockroach@localhost:4132/test?sslmode=disable" \
	go run main.go
