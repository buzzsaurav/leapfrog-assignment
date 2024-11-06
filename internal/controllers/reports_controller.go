// internal/controllers/reports_controller.go

package controllers

import (
	"encoding/json"
	"leapfrog-assignment/internal/services"
	"net/http"
	"strconv"
	"time"
)

func GenerateSalesReport(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")
	productID := r.URL.Query().Get("product_id")
	category := r.URL.Query().Get("category")
	subcategory := r.URL.Query().Get("subcategory") // New subcategory parameter
	location := r.URL.Query().Get("location")

	// Convert dates from strings to time.Time
	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			http.Error(w, "Invalid start_date format", http.StatusBadRequest)
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			http.Error(w, "Invalid end_date format", http.StatusBadRequest)
			return
		}
	}

	// Call the service layer to generate the report
	report, err := services.GenerateSalesReport(startDate, endDate, productID, category, subcategory, location) // Passing subcategory
	if err != nil {
		http.Error(w, "Failed to generate report", http.StatusInternalServerError)
		return
	}

	// Respond with the report
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func GenerateCustomerReport(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")
	minLifetimeValueStr := r.URL.Query().Get("min_lifetime_value")

	// Convert dates from strings to time.Time
	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			http.Error(w, "Invalid start_date format", http.StatusBadRequest)
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			http.Error(w, "Invalid end_date format", http.StatusBadRequest)
			return
		}
	}

	// Convert minLifetimeValue to float64
	var minLifetimeValue float64
	if minLifetimeValueStr != "" {
		minLifetimeValue, err = strconv.ParseFloat(minLifetimeValueStr, 64)
		if err != nil {
			http.Error(w, "Invalid min_lifetime_value format", http.StatusBadRequest)
			return
		}
	}

	// Call the service layer to generate the report
	report, err := services.GenerateCustomerReport(startDate, endDate, minLifetimeValue) // Passing minLifetimeValue as float64
	if err != nil {
		http.Error(w, "Failed to generate report", http.StatusInternalServerError)
		return
	}

	// Respond with the report
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
