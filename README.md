# challenge

This is simple code designing a basic in-memory key-value store in Go. The test are exists in memory folder. Run ```go test``` to execute tests.


## second Question 

For these tables we need to write a query. 

```
Table tasks {
  id integer
  card_id integer
  status sting
  deadline timestamp
  created_at timestamp 
  updated_at timestamp
}

Table cards {
  id integer primary key
  created_at timestamp
  updated_at timestamp
}

Table items {
  id integer
  task_id integer
  name string
  description string
  priority int
  done boolean
  created_at timestamp 
  updated_at timestamp
}
```

1. Sub queries:
```
SELECT * FROM items
WHERE task_id = (SELECT id FROM tasks WHERE id = 1);
```
2. Join
```  
SELECT 
    t.*, 
    c.created_at AS card_created_at, 
    c.updated_at AS card_updated_at,
    i.*
FROM 
    tasks t
LEFT JOIN 
    cards c ON t.card_id = c.id
LEFT JOIN 
    items i ON i.task_id = t.id
```


1. Logical correctness: Sub queries can indeed be the more logically correct way to express certain types of queries, especially when dealing with conditions based on another set of data.

2. Safety: Sub queries can help avoid unintended duplication of results that might occur with joins if not carefully constructed.

3. Performance variability: The performance of sub queries versus joins can vary significantly based on:
   - The specific query optimizer
   - The database management system (DBMS) and its version
   - The particular query structure and data involved

4. Historical context: Joins have traditionally been favored for performance reasons, which has led to the common wisdom that joins are generally better.

5. Evolving optimizers: As query optimizers improve, the performance gap between sub queries and joins is narrowing in many cases.

6. Best practice: Writing queries in a logically coherent way first, then optimizing for performance if necessary, is a sound approach to query development.

[Join vs. sub-query](https://stackoverflow.com/questions/2577174/join-vs-sub-query)