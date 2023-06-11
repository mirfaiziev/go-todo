CREATE TABLE IF NOT EXISTS todo (
  id serial PRIMARY KEY,
  title text NOT NULL,
  state VARCHAR (50)  DEFAULT 'new' NOT NULL
)