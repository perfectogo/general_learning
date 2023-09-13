SELECT 
    name, 
    year
FROM 
    movies
ORDER BY 
    name;

-- Write a new query that retrieves the name, year, and imdb_rating columns of all the movies, ordered highest to lowest by their ratings
SELECT 
    name, 
    year, 
    imdb_rating
FROM 
    movies
ORDER BY 
    imdb_rating DESC;