DO $$ 
BEGIN
SET TIMEZONE='EUROPE/MOSCOW';
CREATE TABLE users (
  id INT PRIMARY KEY,
  reg_time TIMESTAMP WITH TIME ZONE  DEFAULT (now() AT TIME ZONE 'EUROPE/MOSCOW')
);
CREATE TABLE deals(
id SERIAL PRIMARY KEY,
user_id int REFERENCES users(id),
text varchar(1024) not null,
prior NUMERIC(2,0),
status varchar(256) not null,
start_time  TIMESTAMP WITH TIME ZONE,
finish_time  TIMESTAMP WITH TIME ZONE,
treba_povidomiti bool,
bulo_povidomlennia bool
);
END $$;