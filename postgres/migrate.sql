CREATE SCHEMA public;


CREATE TABLE roles (
  id serial NOT NULL,
  name text NOT NULL
);

CREATE TABLE users (
    id serial NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    session_token text NOT NULL,
    role_id int NOT NULL,
    data jsonb NOT NULL
);

ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id);
ALTER TABLE ONLY roles ADD CONSTRAINT roles_pkey PRIMARY KEY (id);
ALTER TABLE ONLY roles ADD CONSTRAINT unique_role UNIQUE (name);

ALTER TABLE ONLY users ADD CONSTRAINT users_roles_fkey FOREIGN KEY (role_id) REFERENCES roles(id);