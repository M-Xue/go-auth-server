CREATE TABLE Users (
	id UUID DEFAULT gen_random_uuid(),
	username varchar(30) UNIQUE NOT NULL,
	email varchar(50) UNIQUE NOT NULL,
	first_name varchar(30) NOT NULL,
	last_name varchar(30) NOT NULL,
	password varchar(30) NOT NULL,
	PRIMARY KEY(id)
);
