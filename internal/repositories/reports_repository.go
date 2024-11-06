// internal/repositories/reports_repository.go

package repositories

import (
	"context"
	"leapfrog-assignment/pkg/database"
	"time"
)

type SalesData struct {
	Customer    string
	Product     string
	Category    string
	Subcategory string
	Region      string
	Quantity    int
	TotalAmount float64
}

func FetchSalesData(startDate, endDate time.Time, productID, category, subcategory, location string) ([]SalesData, error) {
	query := `SELECT c.name AS customer, p.name AS product, cat.name AS category, subcat.name AS subcategory, 
                     c.location AS region, oi.quantity, oi.price * oi.quantity AS total_amount
              FROM orders o
              JOIN customers c ON o.customer_id = c.id
              JOIN order_items oi ON o.id = oi.order_id
              JOIN products p ON oi.product_id = p.id
              JOIN subcategories subcat ON p.subcategory_id = subcat.id
              JOIN categories cat ON subcat.category_id = cat.id
              WHERE (o.order_date BETWEEN $1 AND $2)
              AND (p.id::text = $3 OR $3 = '')
              AND (cat.name = $4 OR $4 = '')
              AND (subcat.name = $5 OR $5 = '')
              AND (c.location = $6 OR $6 = '')`

	rows, err := database.DB.Query(context.Background(), query, startDate, endDate, productID, category, subcategory, location)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []SalesData
	for rows.Next() {
		var data SalesData
		if err := rows.Scan(&data.Customer, &data.Product, &data.Category, &data.Subcategory, &data.Region, &data.Quantity, &data.TotalAmount); err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}

type CustomerData struct {
	Name           string
	LifetimeValue  float64
	OrderFrequency int
}

// FetchCustomerData fetches customer behavior data based on filters
func FetchCustomerData(startDate, endDate time.Time, minLifetimeValue float64) ([]CustomerData, error) {
	query := `SELECT c.name, c.lifetime_value, COUNT(o.id) AS order_frequency
              FROM customers c
              LEFT JOIN orders o ON c.id = o.customer_id AND (o.order_date BETWEEN $1 AND $2 OR $1 IS NULL OR $2 IS NULL)
              WHERE (c.lifetime_value >= $3 OR $3 = 0)
              GROUP BY c.id`

	rows, err := database.DB.Query(context.Background(), query, startDate, endDate, minLifetimeValue)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []CustomerData
	for rows.Next() {
		var data CustomerData
		if err := rows.Scan(&data.Name, &data.LifetimeValue, &data.OrderFrequency); err != nil {
			return nil, err
		}
		results = append(results, data)
	}

	return results, nil
}
