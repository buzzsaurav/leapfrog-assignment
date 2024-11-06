CREATE MATERIALIZED VIEW sales_trends AS
WITH weekly_sales AS (
    SELECT 
        DATE_TRUNC('week', o.order_date) AS week,
        SUM(oi.price * oi.quantity) AS total_sales
    FROM 
        orders o
    JOIN 
        order_items oi ON o.id = oi.order_id
    GROUP BY 
        DATE_TRUNC('week', o.order_date)
),
monthly_sales AS (
    SELECT 
        DATE_TRUNC('month', o.order_date) AS month,
        SUM(oi.price * oi.quantity) AS total_sales
    FROM 
        orders o
    JOIN 
        order_items oi ON o.id = oi.order_id
    GROUP BY 
        DATE_TRUNC('month', o.order_date)
)
SELECT 
    ws.week,
    ws.total_sales AS weekly_sales,
    ms.month,
    ms.total_sales AS monthly_sales
FROM 
    weekly_sales ws
FULL OUTER JOIN 
    monthly_sales ms ON DATE_TRUNC('week', ms.month) = ws.week;
