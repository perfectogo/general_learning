-- Let’s return the name column and a rounded price column.
SELECT 
    name, 
    ROUND(price, 0)
FROM 
    fake_apps;

-- In the last exercise, we were able to get the average price of an app ($2.02365) using this query:
SELECT 
    ROUND(AVG(price), 2)
FROM 
    fake_apps;