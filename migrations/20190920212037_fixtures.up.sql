INSERT INTO countries (code, name) VALUES ('ru', 'Russia');
INSERT INTO countries (code, name) VALUES ('us', 'USA');
INSERT INTO countries (code, name) VALUES ('de', 'Germany');

INSERT INTO users (id, username, email, gender, birthday, country_code) VALUES (1, 'vasya.pupkin', 'vasya.pupkin@example.ru', 'male', '1993-09-15', 'ru');
INSERT INTO users (id, username, email, gender, birthday, country_code) VALUES (2, 'ivan.petrov', 'ivan.petrov@example.ru', 'male', '1978-04-05', 'ru');
INSERT INTO users (id, username, email, gender, birthday, country_code) VALUES (3, 'ann.gorshkova', 'ann.gorshkova@example.ru', 'female', '1988-01-13', 'ru');
INSERT INTO users (id, username, email, gender, birthday, country_code) VALUES (4, 'olga.petrova', 'olga.petrova@example.ru', 'female', '1988-04-13', 'ru');
INSERT INTO users (id, username, email, gender, birthday, country_code) VALUES (5, 'dmitry.barin', 'dmitry.barin@example.ru', 'male', '1988-02-13', 'ru');

ALTER SEQUENCE users_id_seq restart with 6;
