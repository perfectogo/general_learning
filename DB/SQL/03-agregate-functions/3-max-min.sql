-- What is the least number of times an app has been downloaded?
SELECT 
    MIN(downloads)
FROM    
    fake_apps;
 
-- Write a new query that returns the price of the most expensive app.
SELECT 
    MAX(price)
FROM 
    fake_apps;