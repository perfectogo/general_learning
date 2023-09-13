CREATE TABLE customers (
    customer_id	INT     PRIMARY KEY,
    customer_name	    VARCHAR(256),
    address	            VARCHAR(256)
);

CREATE TABLE subscriptions (
    subscription_id	        INT PRIMARY KEY,
    description	            VARCHAR(256),
    price_per_month	        INT,
    subscription_length     TEXT
);


CREATE TABLE orders (
    order_id	        INT PRIMARY KEY,
    customer_id	        INT
    subscription_id 	INT
    purchase_date	    DATE
);


-- First query
SELECT 
    *
FROM 
    orders
JOIN 
    subscriptions 
    ON orders.subscription_id = subscriptions.subscription_id;

-- Second query
SELECT 
    *
FROM 
    orders
JOIN 
    subscriptions 
    ON orders.subscription_id = subscriptions.subscription_id
WHERE 
    subscriptions.description = 'Fashion Magazine';