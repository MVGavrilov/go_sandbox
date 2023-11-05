CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR ( 50 ) NOT NULL,
	last_name   VARCHAR ( 50 ) NOT NULL,
	email       VARCHAR ( 255 ) UNIQUE NOT NULL,
	password    VARCHAR ( 50 ) NOT NULL,
	status      INT,
	created_at  TIMESTAMP NOT NULL
);

CREATE TABLE roles (
   id           SERIAL PRIMARY KEY,
   name         VARCHAR (255) UNIQUE NOT NULL
);

INSERT INTO roles (name) VALUES ('User');
INSERT INTO roles (name) VALUES ('Admin');

CREATE TABLE account_roles (
  user_id       INT NOT NULL,
  role_id       INT NOT NULL,
  grant_date    TIMESTAMP,
  PRIMARY KEY (user_id, role_id),
  FOREIGN KEY (role_id)
      REFERENCES roles (id),
  FOREIGN KEY (user_id)
      REFERENCES users (id)
);

INSERT INTO users (name, last_name, email, password, status, created_at) VALUES ('admin', 'admin', 'admin', 'admin', 0, NOW());
INSERT INTO account_roles (user_id, role_id, grant_date) VALUES (1, 2, NOW());