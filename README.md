Migrate
migrate -path ./migrations -database "postgres://postgres:helloworld@localhost:5432/leapfrog-assignment?sslmode=disable" up

Seed Data
go run cmd/api/main.go -seed


API Endpoints
curl "http://localhost:8080/reports/customers?start_date=2024-01-01&end_date=2024-12-31&min_lifetime_value=100"

curl "http://localhost:8080/reports/sales?start_date=2024-01-01&end_date=2024-12-31&category=Electronics&location=USA"