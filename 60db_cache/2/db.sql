-- b-tree
create index index_name on table (field1, field2);
create index if not exists index_name on table (field1, field2);

-- on big table
create index concurrently if not exists index_name on table (field1, field2);

-- hash
create index concurrently if not exists index_name on table using hash (field1, field2);
select * from users where user_id = 1;

-- gist x1,y1


