SELECT 
    ategory, 
    price,
    AVG(downloads)
FROM 
    fake_apps
GROUP BY 
    category, price;

SELECT 
    category, 
    price,
    AVG(downloads)
FROM 
    fake_apps
GROUP BY 
    1, 2;