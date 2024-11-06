-- Migration to create the products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    subcategory_id INT NOT NULL REFERENCES subcategories(id) ON DELETE SET NULL,
    price DECIMAL(10, 2) NOT NULL,
    UNIQUE(name, subcategory_id) -- Ensures unique product names within each subcategory
);
