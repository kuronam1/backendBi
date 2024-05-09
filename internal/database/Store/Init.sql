CREATE TABLE IF NOT EXISTS groups (
                                    group_id SERIAL PRIMARY KEY,
                                    group_name VARCHAR NOT NULL,
                                    number INTEGER NOT NULL,
                                    speciality VARCHAR NOT NULL,
                                    course INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS users  (
                                      user_id SERIAL PRIMARY KEY,
                                      login VARCHAR UNIQUE NOT NULL,
                                      password VARCHAR NOT NULL,
                                      full_name VARCHAR NOT NULL,
                                      role VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS group_students ( --many to many
                                              group_id INTEGER REFERENCES groups(group_id),
                                              student_id INTEGER REFERENCES users(user_id)
    -- PRIMARY KEY (group_id, student_id)
);


CREATE TABLE IF NOT EXISTS disciplines (
                                           discipline_id SERIAL PRIMARY KEY,
                                           teacher_id INTEGER NOT NULL REFERENCES users(user_id),
                                           discipline_name VARCHAR NOT NULL UNIQUE,
                                           speciality VARCHAR NOT NULL,
                                           course INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS lessons ( --вариант для храниея рассписания всех групп в одной таблице
                                       lesson_id SERIAL PRIMARY KEY,
                                       group_id INTEGER NOT NULL REFERENCES groups(group_id),
                                       time TIMESTAMP NOT NULL,
                                       discipline_id INTEGER NOT NULL REFERENCES disciplines(discipline_id),
                                       audience VARCHAR(10) NOT NULL,
                                       description VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS grades (
                                      grade_id SERIAL PRIMARY KEY,
                                      student_id INTEGER NOT NULL REFERENCES users(user_id),
                                      discipline_id INTEGER NOT NULL REFERENCES disciplines(discipline_id),
                                      grade CHAR(1) CHECK (grade IN ('2', '3', '4', '5', 'н')),
                                      date DATE NOT NULL,
                                      comment TEXT
);

INSERT INTO users (login, password, full_name, role) VALUES ('mylogin', 'mypassword', 'Андрей Горбунов', 'admin')