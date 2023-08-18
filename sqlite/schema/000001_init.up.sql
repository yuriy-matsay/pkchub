CREATE TABLE IF NOT EXISTS сategories (
  id INTEGER NOT NULL PRIMARY KEY,
  name TEXT
);
CREATE TABLE IF NOT EXISTS brands (
  id INTEGER NOT NULL PRIMARY KEY,
  name TEXT
);
CREATE TABLE IF NOT EXISTS models (
  id INTEGER NOT NULL PRIMARY KEY,
  name TEXT,
  image TEXT
);
CREATE TABLE IF NOT EXISTS products (
  id INTEGER NOT NULL PRIMARY KEY,
  idsite INTEGER,
  article TEXT,
  name TEXT,
  model INTEGER,
  brand INTEGER,
  category INTEGER,
  amount INTEGER,
  uah INTEGER,
  usd INTEGER,
  eur INTEGER,
  CONSTRAINT modelRef FOREIGN KEY(model) REFERENCES models (id),
  CONSTRAINT brandRef FOREIGN KEY(brand) REFERENCES brands (id),
  CONSTRAINT categoryRef FOREIGN KEY(category) REFERENCES сategories (id)
);
