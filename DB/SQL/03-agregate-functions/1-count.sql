SELECT 
    COUNT(*) 
FROM 
    fake_apps;

-- Add a WHERE clause in the previous query to count how many free apps are in the table.
SELECT 
  COUNT(*) 
FROM 
  fake_apps
WHERE 
  price = 0;