CREATE TABLE IF NOT EXISTS groups (
                                      group_id SERIAL PRIMARY KEY,
                                      group_name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
                                     user_id SERIAL PRIMARY KEY,
                                     login VARCHAR UNIQUE NOT NULL,
                                     password VARCHAR NOT NULL,
                                     full_name VARCHAR NOT NULL,
                                     role VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS group_users ( --many to many (поменять название)
                                           group_id INTEGER REFERENCES groups(group_id),
                                           user_id INTEGER REFERENCES users(user_id)
);


CREATE TABLE IF NOT EXISTS disciplines (
                                           discipline_id SERIAL PRIMARY KEY,
                                           teacher_id INTEGER NOT NULL REFERENCES users(user_id),
                                           discipline_name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS lessons ( --вариант для храниея рассписания всех групп в одной таблице
                                       lesson_id SERIAL PRIMARY KEY,
                                       group_id INTEGER NOT NULL REFERENCES groups(group_id),
                                       time TIMESTAMP NOT NULL,
                                       discipline_id INTEGER NOT NULL REFERENCES disciplines(discipline_id),
                                       audience VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS grades (
                                      grade_id SERIAL PRIMARY KEY,
                                      student_id INTEGER NOT NULL REFERENCES users(user_id),
                                      discipline_id INTEGER NOT NULL REFERENCES disciplines(discipline_id),
                                      grade CHAR(1) CHECK (grade IN ('2', '3', '4', '5', 'н')),
                                      date DATE NOT NULL,
                                      comment TEXT
);


INSERT INTO groups(group_name) VALUES ('БИ4-1');
INSERT INTO groups(group_name) VALUES ('БИ4-2');

INSERT INTO users (login, password, full_name, role)
    VALUES ('asdfwa', '1234567', 'kolya', 'user(0)');
INSERT INTO users (login, password, full_name, role)
    VALUES ('homemem', 'mypass', 'makson', 'user(1)');
INSERT INTO users (login, password, full_name, role)
    VALUES ('PRICOL', 'qwerty', 'ura', 'user(0)');

INSERT INTO group_users(group_id, user_id) VALUES ('1', '1');
INSERT INTO group_users(group_id, user_id) VALUES ('1', '3');

INSERT INTO disciplines(teacher_id, discipline_name) VALUES ('2', 'Программирование');

INSERT INTO lessons(group_id, time, discipline_id, audience)
    VALUES ('1', CURRENT_TIMESTAMP, '1', '5-309');
INSERT INTO lessons(group_id, time, discipline_id, audience)
    VALUES ('2', CURRENT_TIMESTAMP, '1', '5-310');

INSERT INTO grades(student_id, discipline_id, grade, date, comment)
    VALUES ('1', '1', '4', CURRENT_DATE, 'Контрольрная работа');
INSERT INTO grades(student_id, discipline_id, grade, date, comment)
    VALUES ('3', '1', 'н', CURRENT_DATE, 'Контрольрная работа');

--Вывод оценок кокнретного ученика по user_id
SELECT g.grade_id, g.student_id, g.discipline_id, g.grade, g.date, g.comment
FROM grades g
         INNER JOIN users u on u.user_id = g.grade_id where u.user_id = 1;

--Вывод расписания конкретной группы по user_id
SELECT l.lesson_id, l.group_id, l.time, l.discipline_id, l.audience
FROM lessons l
         INNER JOIN disciplines d on d.discipline_id = l.discipline_id
         INNER JOIN users u on u.user_id = d.discipline_id where u.user_id = 1
ORDER BY l.time;

--Вывод оценок для учителя по группе и конкретной дисциплине
SELECT u.full_name, g.grade, g.date, g.comment
FROM grades g
         JOIN users u ON g.student_id = u.user_id
         JOIN disciplines d ON g.discipline_id = d.discipline_id
         JOIN group_users gu ON u.user_id = gu.user_id
         JOIN groups gr ON gu.group_id = gr.group_id WHERE gr.group_name = 'БИ4-1' AND d.discipline_name = 'Программирование';