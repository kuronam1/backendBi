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
         JOIN group_students gu ON u.user_id = gu.student_id
         JOIN groups gr ON gu.group_id = gr.group_id
WHERE gr.group_name = 'БИ4-1' AND d.discipline_name = 'Программирование';

-- Принадлежность студента к группе
SELECT u.full_name AS student_name, g.group_name
FROM users u
         JOIN group_students gs ON u.user_id = gs.student_id
         JOIN groups g ON gs.group_id = g.group_id
WHERE u.user_id = 12;

-- Вывод расписания для преподавателя
SELECT g.group_name, d.discipline_name, l.time, l.audience, l.description
FROM lessons l
         JOIN disciplines d ON l.discipline_id = d.discipline_id
         JOIN groups g ON l.group_id = g.group_id
WHERE d.teacher_id = 28
ORDER BY l.time;

-- Вывод расписания для студента
SELECT g.group_name, l.time, d.discipline_name, l.audience, l.description, u.full_name AS teacher_name
FROM users u
         JOIN disciplines d ON d.teacher_id = u.user_id
         JOIN lessons l ON l.discipline_id = d.discipline_id
         JOIN groups g ON l.group_id = g.group_id
         JOIN group_students gs ON g.group_id = gs.group_id
WHERE gs.student_id = 1
ORDER BY l.time;

-- Вывод расписания по имени группы
SELECT l.lesson_id, l.time, l.audience, l.description, d.discipline_name
FROM lessons l
        JOIN disciplines d ON l.discipline_id = d.discipline_id
WHERE l.group_id = 1;

INSERT INTO users (login, password, full_name, role) VALUES ('myLogin', 'myPass', 'Андрей Горбунов', 'учитель')
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES (1, 'Программирование', 'ЭВМ', 1)
                                                                          ON CONFLICT DO NOTHING;

SELECT d.discipline_name, g.grade, g.time, g.comment FROM grades g
        JOIN disciplines d ON g.discipline_id = d.discipline_id
        JOIN users u ON g.student_id = u.user_id
WHERE u.user_id = 1

INSERT INTO users (login, password, full_name, role) VALUES ('mylogin', 'mypassword', 'Андрей Горбунов', 'admin');
INSERT INTO users (login, password, full_name, role) VALUES ('sbitneva', 'qweqweqwe', 'Сбитнева Анастасия', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('xzKTO', '473824893', 'Тютюник Николай', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('belov', 'qwerty', 'Белов Ф.Р.', 'teacher');

INSERT INTO groups (group_name, number, speciality, course) VALUES ('БИ','1', 'ЭВМ', '2');
INSERT INTO groups (group_name, number, speciality, course) VALUES ('БИ','2', 'ЭВМ', '2');

INSERT INTO group_students(group_id, student_id) VALUES ('1', '2');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '3');

INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('4', 'Основы философии', 'ЭВМ', '2');

INSERT INTO lessons (group_id, time, discipline_id, audience, description, lesson_order) VALUES ('1', '2024-05-10 8:30:00','1' ,'A-2', 'Лекция', '1');
INSERT INTO lessons (group_id, time, discipline_id, audience, description, lesson_order) VALUES ('2', '2024-05-10 8:30:00','1' ,'A-2', 'Лекция', '1');
INSERT INTO lessons (group_id, time, discipline_id, audience, description, lesson_order) VALUES ('1', '2024-05-11 8:30:00','1' ,'A-2', 'Семенар', '1');
INSERT INTO lessons (group_id, time, discipline_id, audience, description, lesson_order) VALUES ('2', '2024-05-11 10:00:00','1' ,'A-2', 'Семенар', '2');

UPDATE groups SET group_name = 'ЭВМ 2 - 1' WHERE group_id = 1;
UPDATE groups SET group_name = 'ЭВМ 2 - 2' WHERE group_id = 2;

SELECT g.group_name, l.time, d.discipline_name, l.audience, l.description, l.lesson_order
FROM lessons l
         JOIN disciplines d ON l.description = d.discipline_id
         JOIN groups g ON l.group_id = g.group_id
         JOIN users u ON u.user_id = d.teacher_id
WHERE u.full_name = 'Белов Ф.Р.';

SELECT discipline_id, discipline_name FROM disciplines WHERE teacher_id = 4;
SELECT group_id, time, audience, description, lesson_order FROM lessons WHERE discipline_id = 1;
SELECT group_name FROM groups WHERE group_id = 1;

