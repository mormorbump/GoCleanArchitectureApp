CREATE TABLE users (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users(first_name, last_name) VALUES("Patricia", "Smith");
INSERT INTO users(first_name, last_name) VALUES("Linda", "Johnson");
INSERT INTO users(first_name, last_name) VALUES("Mary", "William");
INSERT INTO users(first_name, last_name) VALUES("Robert", "Jones");
INSERT INTO users(first_name, last_name) VALUES("James", "Brown");