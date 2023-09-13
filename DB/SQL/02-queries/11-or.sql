SELECT 
    *
FROM 
    movies
WHERE 
    year > 2014
    OR genre = 'action';

-- Using OR, write a query that returns all movies that are either a romance or a comedy.
SELECT 
    *
FROM 
    movies
WHERE 
    genre = 'romance'
    OR genre = 'comedy';