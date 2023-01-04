BEGIN;

CREATE TABLE people (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255)
);
-- CREATE TABLE taxes(

-- )

INSERT INTO people (name) VALUES
('Владимир'), ('Владислав'), ('Дмитрий');

COMMIT;