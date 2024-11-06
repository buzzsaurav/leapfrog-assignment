// seeders/data_seeder.go

package seeders

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"leapfrog-assignment/pkg/database"
)

// SeedData inserts fake data into all tables
func SeedData() {
	seedCategories()
	seedSubcategories()
	seedCustomers(10)        // Insert 10 customers
	seedProducts(10)         // Insert 10 products
	seedOrders(20)           // Insert 20 orders
	seedOrderItems(50)       // Insert 50 order items
	seedTransactions(20)     // Insert 20 transactions
	fmt.Println("Database seeding complete!")
}

// Helper function to generate a random date
func randomDate() time.Time {
	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	return start.Add(time.Duration(rand.Int63n(int64(end.Sub(start)))))
}

// Seed the Categories table
func seedCategories() {
	categories := []string{"Electronics", "Clothing", "Books"}
	for _, category := range categories {
		_, err := database.DB.Exec(context.Background(),
			`INSERT INTO categories (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`,
			category)
		if err != nil {
			log.Fatalf("Error seeding categories: %v", err)
		}
	}
}

// Seed the Subcategories table
func seedSubcategories() {
	subcategories := map[string][]string{
		"Electronics": {"Mobile Phones", "Laptops", "Cameras"},
		"Clothing":    {"Men's Clothing", "Women's Clothing", "Accessories"},
		"Books":       {"Fiction", "Non-Fiction", "Comics"},
	}

	for category, subs := range subcategories {
		var categoryID int
		err := database.DB.QueryRow(context.Background(),
			`SELECT id FROM categories WHERE name = $1`, category).Scan(&categoryID)
		if err != nil {
			log.Fatalf("Error finding category ID for %s: %v", category, err)
		}

		for _, subcategory := range subs {
			_, err := database.DB.Exec(context.Background(),
				`INSERT INTO subcategories (category_id, name) VALUES ($1, $2) ON CONFLICT DO NOTHING`,
				categoryID, subcategory)
			if err != nil {
				log.Fatalf("Error seeding subcategories for %s: %v", category, err)
			}
		}
	}
}


// Seed the Customers table
func seedCustomers(count int) {
	for i := 0; i < count; i++ {
		_, err := database.DB.Exec(context.Background(),
			`INSERT INTO customers (name, email, signup_date, location, lifetime_value)
			 VALUES ($1, $2, $3, $4, $5)`,
			fmt.Sprintf("Customer %d", i+1),
			fmt.Sprintf("customer%d@example.com", i+1),
			randomDate(),
			[]string{"USA", "Canada", "UK"}[rand.Intn(3)],
			float64(rand.Intn(1000)))
		if err != nil {
			log.Fatalf("Error seeding customers: %v", err)
		}
	}
}

// Seed the Products table
func seedProducts(count int) {
	// Fetch subcategory IDs to associate with products
	var subcategoryIDs []int
	rows, err := database.DB.Query(context.Background(),
		`SELECT id FROM subcategories`)
	if err != nil {
		log.Fatalf("Error fetching subcategory IDs: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatalf("Error scanning subcategory ID: %v", err)
		}
		subcategoryIDs = append(subcategoryIDs, id)
	}

	for i := 0; i < count; i++ {
		subcategoryID := subcategoryIDs[rand.Intn(len(subcategoryIDs))]
		_, err := database.DB.Exec(context.Background(),
			`INSERT INTO products (name, subcategory_id, price) VALUES ($1, $2, $3)`,
			fmt.Sprintf("Product %d", i+1), subcategoryID, float64(rand.Intn(100)+10))
		if err != nil {
			log.Fatalf("Error seeding products: %v", err)
		}
	}
}

// Seed the Orders table
func seedOrders(count int) {
	statuses := []string{"PENDING", "COMPLETED", "CANCELED"}
	for i := 0; i < count; i++ {
		_, err := database.DB.Exec(context.Background(),
			`INSERT INTO orders (customer_id, order_date, status)
			 VALUES ($1, $2, $3)`,
			rand.Intn(10)+1,
			randomDate(),
			statuses[rand.Intn(len(statuses))])
		if err != nil {
			log.Fatalf("Error seeding orders: %v", err)
		}
	}
}

// Seed the Order_Items table
func seedOrderItems(count int) {
    usedCombinations := make(map[string]bool)

    for i := 0; i < count; i++ {
        orderID := rand.Intn(20) + 1
        productID := rand.Intn(10) + 1

        // Check if the combination is already used
        combinationKey := fmt.Sprintf("%d-%d", orderID, productID)
        if usedCombinations[combinationKey] {
            // If the combination is already used, skip this iteration to avoid duplicates
            i--
            continue
        }

        // Mark this combination as used
        usedCombinations[combinationKey] = true

        _, err := database.DB.Exec(context.Background(),
            `INSERT INTO order_items (order_id, product_id, quantity, price)
             VALUES ($1, $2, $3, $4)`,
            orderID,
            productID,
            rand.Intn(5)+1,
            float64(rand.Intn(100)+5))
        if err != nil {
            log.Fatalf("Error seeding order items: %v", err)
        }
    }
}

// Seed the Transactions table
func seedTransactions(count int) {
	statuses := []string{"SUCCESS", "FAILED"}
	for i := 0; i < count; i++ {
		_, err := database.DB.Exec(context.Background(),
			`INSERT INTO transactions (order_id, payment_status, payment_date, total_amount)
			 VALUES ($1, $2, $3, $4)`,
			rand.Intn(20)+1,
			statuses[rand.Intn(len(statuses))],
			randomDate(),
			float64(rand.Intn(500)+50))
		if err != nil {
			log.Fatalf("Error seeding transactions: %v", err)
		}
	}
}
