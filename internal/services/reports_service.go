// internal/services/reports_service.go

package services

import (
    "leapfrog-assignment/internal/repositories"
    "fmt"
    "time"
)

type SalesReport struct {
    TotalSales         int                      `json:"total_sales"`
    AverageOrderValue  float64                  `json:"average_order_value"`
    ProductsSold       int                      `json:"products_sold"`
    RevenueByCustomer  map[string]float64       `json:"revenue_by_customer"`
    RevenueByProduct   map[string]float64       `json:"revenue_by_product"`
    RevenueByRegion    map[string]float64       `json:"revenue_by_region"`
}

func GenerateSalesReport(startDate, endDate time.Time, productID, category, subcategory, location string) (SalesReport, error) {
    // Fetch sales data with applied filters from repository
    data, err := repositories.FetchSalesData(startDate, endDate, productID, category, subcategory, location)
    if err != nil {
        return SalesReport{}, err
    }

    // Aggregate results
    report := SalesReport{
        RevenueByCustomer: make(map[string]float64),
        RevenueByProduct:  make(map[string]float64),
        RevenueByRegion:   make(map[string]float64),
    }

    // Process data to calculate totals and averages
    for _, record := range data {
        report.TotalSales++
        report.ProductsSold += record.Quantity
        report.AverageOrderValue += record.TotalAmount
        report.RevenueByCustomer[record.Customer] += record.TotalAmount
        report.RevenueByProduct[record.Product] += record.TotalAmount
        report.RevenueByRegion[record.Region] += record.TotalAmount
    }

    // Finalize average calculation
    if report.TotalSales > 0 {
        report.AverageOrderValue /= float64(report.TotalSales)
    }

    return report, nil
}

type CustomerReport struct {
    TotalCustomers         int                      `json:"total_customers"`
    AverageOrderFrequency  float64                  `json:"average_order_frequency"`
    TotalLifetimeValue     float64                  `json:"total_lifetime_value"`
    CustomerSegments       map[string]int           `json:"customer_segments"`
}

func GenerateCustomerReport(startDate, endDate time.Time, minLifetimeValue float64) (CustomerReport, error) {
    // Fetch customer data with applied filters from the repository
    data, err := repositories.FetchCustomerData(startDate, endDate, minLifetimeValue)
    if err != nil {
        return CustomerReport{}, fmt.Errorf("failed to fetch customer data: %w", err)
    }

    // Aggregate results
    report := CustomerReport{
        CustomerSegments: make(map[string]int),
    }

    // Process data to calculate totals and averages
    for _, record := range data {
        report.TotalCustomers++
        report.TotalLifetimeValue += record.LifetimeValue

        // Assuming orderFrequency is calculated based on orders
        report.AverageOrderFrequency += float64(record.OrderFrequency)

        // Segmentation based on lifetime value
        if record.LifetimeValue < 100 {
            report.CustomerSegments["Low Value"]++
        } else if record.LifetimeValue >= 100 && record.LifetimeValue <= 500 {
            report.CustomerSegments["Medium Value"]++
        } else {
            report.CustomerSegments["High Value"]++
        }
    }

    // Finalize average calculation
    if report.TotalCustomers > 0 {
        report.AverageOrderFrequency /= float64(report.TotalCustomers)
    }

    return report, nil
}
