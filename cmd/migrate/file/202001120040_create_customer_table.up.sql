CREATE TABLE IF NOT EXISTS customers(
   id serial PRIMARY KEY,
   name VARCHAR (50),
   password text NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL
);
