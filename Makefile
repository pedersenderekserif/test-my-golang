dependencies:
	go mod tidy && \
	go mod download

test:
	rm -f test.log && \
	rm -f test.html && \
	rm -f test.out && \
	go test ./... -covermode=count -v -short -coverprofile test.out && \
	go tool cover -html=test.out -o test.html && \
	go tool cover -func=test.out

build: dependencies
	rm -rf bin/test-my-golang && \
	go build -o bin/test-my-golang

swagger:
	rm -f .swagger/swagger.json && \
	go generate && \
	swagger version && \
	swagger validate .swagger/swagger.json

db-docker:
	docker-compose up -d --remove-orphans
	sleep 5

db-copy:
	psql "dbname='teams' user='teams' password='password' host='localhost' port='4532'" -f .db/scripts/copy_data.sql

db-seed:
	psql "dbname='teams' user='teams' password='password' host='localhost' port='4532'" -f .db/scripts/insert_data.sql

db-migrate:
	migrate -source file://.db/schema -database postgres://teams:password@localhost:4532/teams?sslmode=disable up 

db: db-docker db-migrate db-seed