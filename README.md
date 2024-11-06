## Note to the evaluator
a. The assignment was completed in under 2 days, following the Tihar holidays.\
b. This is the first time I am using Golang. I am yet to familiarize myself with the best practices of the language. For this project, I implemented a simple controller-service-repository architecture.

## Migrate

```http
migrate -path ./migrations -database "postgres://postgres:helloworld@localhost:5432/leapfrog-assignment?sslmode=disable" up
````

## Seed Data

```http
go run cmd/api/main.go -seed
````

## SQL Reference

SQL queries for the assignment are located in assignment-sql-commands folder in the root directory.

a. Customer Retention Calculation.sql\
b. Ranking Customers by Total Spending.sql\
c. Summarize Sales per Category Including Subcategories.sql\
d. Sales Trends Over Different Time Periods.sql\
e. Materialized View Example.sql


## API Reference

#### Get Sales Report

```http
  GET /reports/sales
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `start_date` | `string` | **Required** (2024-01-01)|
| `end_date` | `string` | **Required** (2024-12-31)|
| `category` | `string` | Electronics, Clothing, Books|
| `location` | `string` | USA, Canada, UK|


#### Get Customer Lifetime Value

```http
  GET /reports/customers
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `start_date` | `string` | **Required** (2024-01-01)|
| `end_date` | `string` | **Required** (2024-12-31)|
| `min_lifetime_value` | `string` | 100|

#### Example

curl "http://localhost:8080/reports/customers?start_date=2024-01-01&end_date=2024-12-31&min_lifetime_value=100"

curl "http://localhost:8080/reports/sales?start_date=2024-01-01&end_date=2024-12-31&category=Electronics&location=USA"

