CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    title text NOT NULL,
    state varchar (50) NOT NULL DEFAULT 'new'
)