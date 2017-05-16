insert into store_statuses
select p.id, b.bookstore_id,  Round(RAND() * 20, 0) as stack
, Round(RAND() * 300, 0)as stock,
now() as create_at,
now() as update_at,
NULL as delete_at
from products p
cross join book_stores b;

insert into sales_statuses
select p.id, b.bookstore_id,  Round(RAND() * 15, 0) as day
, Round(RAND() * 60, 0)as week
, Round(RAND() * 200, 0) as month,
now() as create_at,
now() as update_at,
NULL as delete_at
from products p
cross join book_stores b