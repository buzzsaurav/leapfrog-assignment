WITH category_sales AS (
    SELECT 
        c.name AS category_name,
        sc.name AS subcategory_name,
        SUM(oi.price * oi.quantity) AS total_sales
    FROM 
        categories c
    JOIN 
        subcategories sc ON c.id = sc.category_id
    JOIN 
        products p ON p.subcategory_id = sc.id
    JOIN 
        order_items oi ON p.id = oi.product_id
    JOIN 
        orders o ON oi.order_id = o.id
    GROUP BY 
        c.id, sc.id
)
SELECT 
    cs.category_name,
    cs.subcategory_name,
    COALESCE(cs.total_sales, 0) AS total_sales
FROM 
    category_sales cs
ORDER BY 
    cs.category_name, cs.subcategory_name;
