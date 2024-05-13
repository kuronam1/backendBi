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

INSERT INTO users (login, password, full_name, role) VALUES ('mylogin', 'mypassword', 'Андрей Горбунов', 'admin');
INSERT INTO users (login, password, full_name, role) VALUES ('aliev', '1111111111', 'Алиев Алексей Алексеевич', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('boichenco', '22222222', 'Бойченко Александр Евгеньевич', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('zarubin', '3333333333', 'Зарубин Николай Владимирович', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('zizichenko', '4444444444', 'Зинченко Ольга Александровна', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('zolotorev', '555555555', 'Золоторёв Александр Владимирович', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('corovkin', '666666666', 'Коровкин Игорь Алексеевич', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('kyznech', '777777777', 'Кузнецов Ярослав Олегович', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('daniil', '8888888888', 'Львутин Даниил Владимирович', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('maxan', '9999999999', 'Маханов Юрий Игоревич', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('navalski', '1010101010', 'Навальский Александр Сергеевич', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('popyan', '1212121212', 'Полуян Алексей Алексеевич', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('salimov', '1313131313', 'Салимов Тимур Саламатович', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('scotnicov', '1414141414', 'Скотников Максим Кириллович', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('tkachenco', '1515151515', 'Ткаченко Анна Кирилловна', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('fedor', '1616161616', 'Фадеев Федор Денисович', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('serega', '1717171717', 'Харькин Сергей Михайлович', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('anna', '1818181818', 'Анна Щербинина Павловна', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('shewkai', '1919191919', 'Щуцкая Наталья Николаевна', 'student');

INSERT INTO users (login, password, full_name, role) VALUES ('belov', 'qwerty', 'Белов Ф.Р.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('ballechene', 'asdfwa', 'Бальчюнене И.П.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('borodkin', '131231', 'Бородкин Е.Ю.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('gysarova', '53454234', 'Гусарова Т.В.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('xaystova', '823188132', 'Хаустова Н.Е.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('jilina', '0848173', 'Жилина Т.А.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('knyazev', '31623694', 'Князев С.Л.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('losev', '89492392', 'Лосев А.Н.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('pivtorawkaya', '121324435', 'Пивторацкая Н.И.', 'teacher');
INSERT INTO users (login, password, full_name, role) VALUES ('dybanova', '84349832', 'Дубанова О.В.', 'teacher');

INSERT INTO users (login, password, full_name, role) VALUES ('chimee', '1111111111', 'Муравьев Роман Матвеевич', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('xhosaly', '22222222', 'Никитин Кирилл Дмитриевич', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('violan', '3333333333', 'Жилин Дмитрий Денисович', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('ellahel', '4444444444', 'Коновалов Роман Сергеевич', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('arillys', '555555555', 'Грачев Александр Алексеевич', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('mokermi', '666666666', 'Щербакова Анастасия Михайловна', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('prycee', '777777777', 'Малышев Максим Кириллович', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('wendol', '8888888888', 'Трошин Матвей Билалович', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('nganael', '9999999999', 'Егорова Дарья Алексеевна', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('amlisi', '1010101010', 'Александров Алексей Дмитриевич', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('yridgi', '1212121212', 'Фокина Есения Дмитриевна', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('justia', '1313131313', 'Бобров Тимур Иванович', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('maniemu', '1414141414', 'Киреева Арина Степановна', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('eikiton', '1515151515', 'Беляков Степан Алексеевич', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('blythe', '1616161616', 'Тихонова Амира Богдановна', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('bianne', '1717171717', 'Широкова Диана Денисовна', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('stinah', '1818181818', 'Чернова Ульяна Георгиевна', 'parent');
INSERT INTO users (login, password, full_name, role) VALUES ('madgett', '1919191919', 'Русаков Тимофей Викторович', 'parent');

INSERT INTO groups (group_name, number, speciality, course) VALUES ('ИСП','1', 'Информационные системы и программирование', '2');
INSERT INTO groups (group_name, number, speciality, course) VALUES ('ИСП','2', 'Информационные системы и программирование', '2');

INSERT INTO group_students(group_id, student_id) VALUES ('1', '2');
INSERT INTO group_students(group_id, student_id) VALUES ('1', '4');
INSERT INTO group_students(group_id, student_id) VALUES ('1', '6');
INSERT INTO group_students(group_id, student_id) VALUES ('1', '8');
INSERT INTO group_students(group_id, student_id) VALUES ('1', '10');
INSERT INTO group_students(group_id, student_id) VALUES ('1', '12');
INSERT INTO group_students(group_id, student_id) VALUES ('1', '14');
INSERT INTO group_students(group_id, student_id) VALUES ('1', '16');
INSERT INTO group_students(group_id, student_id) VALUES ('1', '18');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '3');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '5');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '7');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '9');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '11');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '13');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '15');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '17');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '19');

-----------------------------------------------------------------

INSERT INTO users (login, password, full_name, role) VALUES ('mylogin', 'mypassword', 'Андрей Горбунов', 'admin');
INSERT INTO users (login, password, full_name, role) VALUES ('sbitneva', 'qweqweqwe', 'Сбитнева Анастасия', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('xzKTO', '473824893', 'Тютюник Николай', 'student');
INSERT INTO users (login, password, full_name, role) VALUES ('belov', 'qwerty', 'Белов Ф.Р.', 'teacher');

INSERT INTO groups (group_name, number, speciality, course) VALUES ('ЭВМ 2 - 1','1', 'ЭВМ', '2');
INSERT INTO groups (group_name, number, speciality, course) VALUES ('ЭВМ 2 - 2','2', 'ЭВМ', '2');

INSERT INTO group_students(group_id, student_id) VALUES ('1', '2');
INSERT INTO group_students(group_id, student_id) VALUES ('2', '3');

INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('4', 'Основы философии', 'ЭВМ', '2');

INSERT INTO lessons (group_id, time, discipline_id, teacher_id, audience, description, lesson_order) VALUES ('1', '2024-05-10 8:30:00','1', 4 ,'A-2', 'Лекция', '1');
INSERT INTO lessons (group_id, time, discipline_id, teacher_id, audience, description, lesson_order) VALUES ('2', '2024-05-10 8:30:00','1', 4 ,'A-2', 'Лекция', '1');
INSERT INTO lessons (group_id, time, discipline_id, teacher_id, audience, description, lesson_order) VALUES ('1', '2024-05-11 8:30:00','1', 4 ,'A-2', 'Семенар', '1');
INSERT INTO lessons (group_id, time, discipline_id, teacher_id, audience, description, lesson_order) VALUES ('2', '2024-05-11 10:00:00','1', 4 ,'A-2', 'Семенар', '2');

-----------------------------------------------------------------

UPDATE groups SET group_name = 'ЭВМ 2 - 1' WHERE group_id = 1;
UPDATE groups SET group_name = 'ЭВМ 2 - 2' WHERE group_id = 2;

SELECT d.discipline_name, l.audience, l.description, l.lesson_order
FROM lessons l
         JOIN disciplines d ON l.description = d.discipline_id
         JOIN users u ON u.user_id = d.teacher_id
WHERE u.user_id = ;

SELECT discipline_id, discipline_name FROM disciplines WHERE teacher_id = 4;
SELECT group_id, time, audience, description, lesson_order FROM lessons WHERE discipline_id = 1;
SELECT group_name FROM groups WHERE group_id = 1;

SELECT group_id, speciality, group_name, number, course FROM groups WHERE group_name = 'ЭВМ2-2';
INSERT INTO grades VALUES (1, 2, 1, 4, '02-01-2024', 'Контрольная работа');

SELECT g.grade_id, g.grade, g.date, g.comment, u.full_name
FROM grades g
         JOIN users u ON g.student_id = u.user_id
         JOIN disciplines d ON d.discipline_id = g.discipline_id
         JOIN groups gr ON d.speciality = gr.speciality AND d.course = gr.course
WHERE gr.group_name = 'ЭВМ2-2' AND d.discipline_name = 'Основы философии';

SELECT g.group_id, g.speciality, g.group_name, g.number, g.course
FROM groups g
        JOIN group_students gs ON g.group_id = gs.group_id
WHERE gs.student_id = 2;

SELECT g.group_id, group_name, number, speciality, course
FROM groups g
        JOIN group_students gr ON g.group_id = gr.group_id WHERE student_id = 2;

SELECT g.grade_id, g.grade, g.date, g.comment, u.full_name
FROM grades g
         JOIN users u ON g.student_id = u.user_id
         JOIN disciplines d ON d.discipline_id = g.discipline_id
         JOIN groups gr ON d.speciality = gr.speciality AND d.course = gr.course
WHERE gr.group_name = 'ЭВМ2-2' AND d.discipline_name = 'Основы философии';

SELECT group_id, speciality, group_name, number, course FROM groups WHERE group_name = 'ЭВМ2-2';

INSERT INTO grades VALUES (1, 2, 1, '4', '11-05-2024', 'Контрольная работа');
INSERT INTO grades VALUES (2, 3, 1, '3', '10-05-2024', 'Контрольная работа');
UPDATE groups SET group_name = 'ЭВМ2-2' WHERE group_id = 2;