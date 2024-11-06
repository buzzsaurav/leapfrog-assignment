// main.go

package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "leapfrog-assignment/configs"
    "leapfrog-assignment/internal/controllers"
    "leapfrog-assignment/pkg/database"
    "leapfrog-assignment/seeders"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    cfg := configs.LoadConfig()
    database.ConnectDB(cfg)

    // Check if we should seed data
    seed := flag.Bool("seed", false, "Seed the database with test data")
    flag.Parse()

    if *seed {
        fmt.Println("Seeding the database with test data...")
        seeders.SeedData()
        fmt.Println("Database seeding complete.")
        return
    }

    r := mux.NewRouter()

    // Define your routes
    r.HandleFunc("/reports/sales", controllers.GenerateSalesReport).Methods("GET")
	r.HandleFunc("/reports/customers", controllers.GenerateCustomerReport).Methods("GET")


    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
