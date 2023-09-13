SELECT 
    price, 
    COUNT(*) 
FROM 
    fake_apps
GROUP BY 
    price;

SELECT 
    price, 
    COUNT(*) 
FROM 
    fake_apps
WHERE 
    downloads > 20000
GROUP BY 
    price;

SELECT 
    category, 
    SUM(downloads)
FROM 
    fake_apps
GROUP BY 
    category;
 
