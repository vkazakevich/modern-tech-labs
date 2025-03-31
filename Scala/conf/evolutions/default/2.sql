# categories schema

# --- !Ups
create table categories (
  `id` integer not null primary key autoincrement,
  `name` varchar(255) not null
)

# --- !Downs
drop table categories; 