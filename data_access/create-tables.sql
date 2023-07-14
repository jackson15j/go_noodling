/* https://go.dev/doc/tutorial/database-access */
DROP TABLE IF EXISTS data_access.album;
DROP SCHEMA IF EXISTS data_access;
/* https://www.postgresql.org/docs/current/sql-createschema.html */
CREATE SCHEMA data_access;
/* https://www.postgresql.org/docs/current/sql-createtable.html */
CREATE TABLE data_access.album (
  id         SERIAL primary key,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL
);

INSERT INTO data_access.album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);

/* Verification */
SELECT * FROM data_access.album;
