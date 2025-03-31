# cart schema

# --- !Ups
create table cart (
  `id` integer not null primary key autoincrement,
  `product_id` integer not null,
  `quantity` integer not null
)

# --- !Downs
drop table cart; 