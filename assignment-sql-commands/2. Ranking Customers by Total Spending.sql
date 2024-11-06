WITH customer_spending AS (
    SELECT 
        c.id AS customer_id,
        c.name AS customer_name,
        c.location,
        SUM(oi.price * oi.quantity) AS total_spending
    FROM 
        customers c
    JOIN 
        orders o ON c.id = o.customer_id
    JOIN 
        order_items oi ON o.id = oi.order_id
    GROUP BY 
        c.id, c.name, c.location
),
ranked_customers AS (
    SELECT 
        customer_id,
        customer_name,
        location,
        total_spending,
        RANK() OVER (PARTITION BY location ORDER BY total_spending DESC) AS spending_rank
    FROM 
        customer_spending
)
SELECT 
    customer_id,
    customer_name,
    location,
    total_spending,
    spending_rank
FROM 
    ranked_customers
WHERE 
    spending_rank <= 10  -- Use a parameter to specify the top N customers
ORDER BY 
    location, spending_rank;
