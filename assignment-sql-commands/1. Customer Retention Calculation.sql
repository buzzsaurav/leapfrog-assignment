WITH first_purchase AS (
    SELECT 
        customer_id,
        MIN(order_date) AS first_purchase_date
    FROM 
        orders
    GROUP BY 
        customer_id
),
customer_activity AS (
    SELECT 
        fp.customer_id,
        COUNT(CASE WHEN o.order_date >= fp.first_purchase_date + INTERVAL '3 months' THEN 1 END) AS retained_after_3_months,
        COUNT(CASE WHEN o.order_date >= fp.first_purchase_date + INTERVAL '6 months' THEN 1 END) AS retained_after_6_months,
        COUNT(CASE WHEN o.order_date >= fp.first_purchase_date + INTERVAL '12 months' THEN 1 END) AS retained_after_12_months
    FROM 
        first_purchase fp
    LEFT JOIN 
        orders o ON fp.customer_id = o.customer_id
    GROUP BY 
        fp.customer_id
)
SELECT 
    COUNT(DISTINCT fp.customer_id) AS total_customers,
    SUM(retained_after_3_months) AS retained_after_3_months,
    SUM(retained_after_6_months) AS retained_after_6_months,
    SUM(retained_after_12_months) AS retained_after_12_months
FROM 
    customer_activity
JOIN 
    first_purchase fp ON customer_activity.customer_id = fp.customer_id;
