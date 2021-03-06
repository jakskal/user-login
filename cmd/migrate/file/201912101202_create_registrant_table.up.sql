CREATE TABLE IF NOT EXISTS registrants(
   id serial PRIMARY KEY,
   name VARCHAR (50),
   password text NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   activation_code VARCHAR(6) NOT NULL,
   is_activated boolean
);
