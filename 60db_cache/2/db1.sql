-- MB вся таблица
SELECT pg_size_pretty(pg_total_relation_size('table1'));

-- данные
SELECT pg_size_pretty(pg_table_size('table2'));

-- индексы
SELECT pg_size_pretty(pg_indexes_size('table3'));

