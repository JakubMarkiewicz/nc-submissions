-- Write your query below
SELECT name from customers LEFT JOIN orders ON customers.id = orders.customer_id where orders.id is null;