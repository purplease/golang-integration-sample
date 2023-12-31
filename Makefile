build:
	docker build -t golang-integration-sample .

run:
	docker-compose down
	docker-compose up

reload:
	docker-compose down
	docker image rm golang-integration-sample
	docker image rm golang-integration-sample_app
	docker build -t golang-integration-sample .
	docker-compose up

test:
	go test -count=1 -failfast -v -race ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html && go tool cover -func coverage.out

generate-mocks:
	mockgen -source=pkg/infrastructure/repository/diner.go -destination=pkg/infrastructure/mocks/repository/diner.go -package mocks
	mockgen -source=pkg/infrastructure/repository/menu.go -destination=pkg/infrastructure/mocks/repository/menu.go -package mocks
	mockgen -source=pkg/infrastructure/repository/order.go -destination=pkg/infrastructure/mocks/repository/order.go -package mocks

generate:
	swag init -g pkg/infrastructure/rest/routes/routes.go

db-update:
	docker run golang-integration-sample update-db

	
