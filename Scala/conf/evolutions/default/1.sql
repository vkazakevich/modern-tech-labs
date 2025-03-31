# products schema

# --- !Ups
create table products (
  `id` integer not null primary key autoincrement,
  `title` varchar(255) not null,
  `description` TEXT null,
  `price` NUMERIC not null
)

# --- !Downs
drop table products; 
