CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   name VARCHAR (50),
   password text NOT NULL,
   role VARCHAR(10) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL
);
