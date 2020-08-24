test-pre:
	docker run --rm -d --name chassis-mysql-ut -e MYSQL_ALLOW_EMPTY_PASSWORD=true -e MYSQL_DATABASE=test -p 3306:3306 mysql:8
	docker run --rm -d --name chassis-postgre-ut -e POSTGRES_DB=postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=test -p 5432:5432 postgres:10
	docker run --rm -d --name chassis-redis-ut -p 6379:6379 redis:5.0
test:
	PG_CONF_FILE=${PWD}/configs/app.yml go test -v ./...

test-post:
	docker stop chassis-mysql-ut
	docker stop chassis-postgre-ut
	docker stop chassis-redis-ut

