CREATE TABLE IF NOT EXISTS specialities (
                                            speciality_id SERIAL PRIMARY KEY,
                                            speciality_name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS groups (
                                      group_id SERIAL PRIMARY KEY,
                                      group_name VARCHAR NOT NULL,
                                      number INTEGER NOT NULL,
                                      speciality VARCHAR NOT NULL REFERENCES specialities(speciality_name),
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
                                              student_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE
    -- PRIMARY KEY (group_id, student_id)
);


CREATE TABLE IF NOT EXISTS disciplines (
                                           discipline_id SERIAL PRIMARY KEY,
                                           teacher_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
                                           discipline_name VARCHAR,
                                           speciality VARCHAR REFERENCES specialities(speciality_name),
                                           course INTEGER
);

CREATE TABLE IF NOT EXISTS lessons ( --вариант для храниея рассписания всех групп в одной таблице
                                       lesson_id SERIAL PRIMARY KEY,
                                       group_id INTEGER NOT NULL REFERENCES groups(group_id),
                                       time DATE NOT NULL,
                                       discipline_id INTEGER NOT NULL REFERENCES disciplines(discipline_id),
                                       teacher_id INTEGER NOT NULL REFERENCES users(user_id),
                                       audience VARCHAR(10) NOT NULL,
                                       description VARCHAR NOT NULL,
                                       subject VARCHAR DEFAULT NULL,
                                       homework VARCHAR DEFAULT NULL,
                                       lesson_order INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS grades (
                                      grade_id SERIAL PRIMARY KEY,
                                      student_id INTEGER NOT NULL REFERENCES users(user_id),
                                      discipline_id INTEGER NOT NULL REFERENCES disciplines(discipline_id),
                                      grade CHAR(1) CHECK (grade IN ('2', '3', '4', '5', 'н')),
                                      date DATE NOT NULL,
                                      comment TEXT
);

CREATE TABLE IF NOT EXISTS parent_students (
                                               parent_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
                                               student_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE
);

INSERT INTO specialities (speciality_name) VALUES ('Специальные машины и устройства');
INSERT INTO specialities (speciality_name) VALUES ('Автоматические системы управления');
INSERT INTO specialities (speciality_name) VALUES ('Компьютерные системы и комплексы');
INSERT INTO specialities (speciality_name) VALUES ('Контроль работы измерительных приборов');
INSERT INTO specialities (speciality_name) VALUES ('Экономика и бухгалтерский учет');
INSERT INTO specialities (speciality_name) VALUES ('Технология металлообрабатывающего производства');
INSERT INTO specialities (speciality_name) VALUES ('Технология машиностроения');
INSERT INTO specialities (speciality_name) VALUES ('Информационные системы и программирование');
INSERT INTO specialities (speciality_name) VALUES ('Техническая эксплуатация летательных аппаратов и двигателей');
INSERT INTO specialities (speciality_name) VALUES ('Техническая эксплуатация электрифицированных и пилотажно-навигационных комплексов');

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
INSERT INTO users
(login, password, full_name, role) VALUES ('arillys', '555555555', 'Грачев Александр Алексеевич', 'parent');
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

INSERT INTO parent_students(parent_id, student_id) VALUES ('30', '2');
INSERT INTO parent_students(parent_id, student_id) VALUES ('31', '3');
INSERT INTO parent_students(parent_id, student_id) VALUES ('32', '4');
INSERT INTO parent_students(parent_id, student_id) VALUES ('33', '5');
INSERT INTO parent_students(parent_id, student_id) VALUES ('34', '6');
INSERT INTO parent_students(parent_id, student_id) VALUES ('35', '7');
INSERT INTO parent_students(parent_id, student_id) VALUES ('36', '8');
INSERT INTO parent_students(parent_id, student_id) VALUES ('37', '9');
INSERT INTO parent_students(parent_id, student_id) VALUES ('38', '10');
INSERT INTO parent_students(parent_id, student_id) VALUES ('39', '11');
INSERT INTO parent_students(parent_id, student_id) VALUES ('40', '12');
INSERT INTO parent_students(parent_id, student_id) VALUES ('41', '13');
INSERT INTO parent_students(parent_id, student_id) VALUES ('42', '14');
INSERT INTO parent_students(parent_id, student_id) VALUES ('43', '15');
INSERT INTO parent_students(parent_id, student_id) VALUES ('44', '16');
INSERT INTO parent_students(parent_id, student_id) VALUES ('45', '17');
INSERT INTO parent_students(parent_id, student_id) VALUES ('46', '18');
INSERT INTO parent_students(parent_id, student_id) VALUES ('47', '19');

INSERT INTO groups (group_name, number, speciality, course) VALUES ('ИСП21','1', 'Информационные системы и программирование', '2');
INSERT INTO groups (group_name, number, speciality, course) VALUES ('ИСП22','2', 'Информационные системы и программирование', '2');

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

INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('20', 'Основы философии', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('21', 'Иностранный язык в профильной деятельности', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('22', 'Физическая культура', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('23', 'Теория вероятностей и математическая статистика', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('24', 'Архитектура аппаратных средств',
                                                                                  'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('25', 'Основы алгоритмизации и программировая', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('26', 'Основы проектирования баз данных', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('27', 'Компьютерные сети', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('28', 'Внедрение и поддержка программного обеспечения компьютерных систем', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('29', 'Обеспечение качества функционирования компьютерных систем', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('28', 'Учебная практика', 'Информационные системы и программирование', '2');
INSERT INTO disciplines (teacher_id, discipline_name, speciality, course) VALUES ('30', 'Производственная практика', 'Информационные системы и программирование', '2');