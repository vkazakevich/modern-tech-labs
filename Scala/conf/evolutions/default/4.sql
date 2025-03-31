# Sample data for categories schema

# --- !Ups

INSERT INTO categories (name)
VALUES ('Coffee'), ('Tee'), ('Water');

# --- !Downs

DELETE FROM categories
WHERE name IN (
    'Coffee',
    'Tee',
    'Water'
);

# Sample data for products schema

# --- !Ups

INSERT INTO products (title, price)
VALUES ('Product 1', 12.5), ('Product 2', 25.95), ('Product 3', 99);

# --- !Downs

DELETE FROM products
WHERE title IN (
    'Product 1',
    'Product 2',
    'Product 3'
);

# Sample data for cart schema

# --- !Ups

INSERT INTO cart (product_id, quantity)
VALUES (1, 2), (2, 1), (3, 99);

# --- !Downs

DELETE FROM products
WHERE title IN (
    'Product 1',
    'Product 2',
    'Product 3'
);