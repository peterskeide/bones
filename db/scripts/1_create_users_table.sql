CREATE TABLE IF NOT EXISTS users
(
 	id serial NOT NULL,
 	password character varying(255) NOT NULL,
 	email character varying(255) NOT NULL,
 	CONSTRAINT users_pkey PRIMARY KEY (id)
);
