-- как инспектить запросы?
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   age INT NOT NULL,
   city VARCHAR(255) NOT NULL
);

-- начинаем explain всегда с транзакции
begin;
EXPLAIN ANALYZE
SELECT * FROM users
WHERE age > 30
ORDER BY name;
rollback;

-- output
/*
Sort  (cost=40.92..42.46 rows=614 width=40) (actual time=0.147..0.162 rows=35 loops=1)
  Sort Key: name
  Sort Method: quicksort  Memory: 29kB
  ->  Seq Scan on users  (cost=0.00..17.67 rows=614 width=40) (actual time=0.031..0.090 rows=35 loops=1)
        Filter: (age > 30)
        Rows Removed by Filter: 230
Planning Time: 0.098 ms
Execution Time: 0.229 ms
  */
