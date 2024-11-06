// internal/handlers/user.go
package handlers

import (
    "encoding/json"
    "net/http"
    "log"
    "leapfrog-assignment/internal/models"
    "leapfrog-assignment/pkg/database"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User

    rows, err := database.DB.Query(r.Context(), "SELECT id, name, email FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user) 
    }

    // Check for any errors after looping through rows
    if err = rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(users); err != nil {
        log.Printf("Error encoding response to JSON: %v", err)
        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
    }
}
