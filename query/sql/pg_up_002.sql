INSERT INTO equipments (equip) VALUES ('Лабораторное оборудование');
INSERT INTO equipments (equip) VALUES ('Практическое оборудование');
INSERT INTO equipments (equip) VALUES ('Лекционное оборудование');

INSERT INTO rooms (number) VALUES ('301а');
INSERT INTO rooms (number) VALUES ('104');
INSERT INTO rooms (number) VALUES ('263б');

INSERT INTO courses(naming, description) VALUES ('биоинженерия', 'описание предмета?');

INSERT INTO specialities (naming, code, depart_id) VALUES ('Информационные системы и технологии', '9.1.3', 1);
INSERT INTO specialities (naming, code, depart_id) VALUES ('Разработка программных решений', '9.1.4', 1);
INSERT INTO specialities (naming, code, depart_id) VALUES ('Программирование на графических процессорах', '9.1.5', 1);

INSERT INTO groups (depart_id, spec_id) VALUES (1, 1);
INSERT INTO groups (depart_id, spec_id) VALUES (1, 2);
INSERT INTO groups (depart_id, spec_id) VALUES (1, 3);

INSERT INTO students (full_name, group_id) VALUES ('Владимир Витальевич Теренко', 1);
INSERT INTO students (full_name, group_id) VALUES ('Владислав Сергеевич Филлипов', 1);
INSERT INTO students (full_name, group_id) VALUES ('Виктор Николаевич Сорокин', 1);
INSERT INTO students (full_name, group_id) VALUES ('Вячеслав Викторович Фидзе', 1);
INSERT INTO students (full_name, group_id) VALUES ('Евгения Григорьевна Павленко', 1);
INSERT INTO students (full_name, group_id) VALUES ('Татьяна Валерьевна Головчак', 1);

INSERT INTO students (full_name, group_id) VALUES ('Дмитрий Викторович Денькин', 2);
INSERT INTO students (full_name, group_id) VALUES ('Оксана Евгеньевна Лапшина', 2);
INSERT INTO students (full_name, group_id) VALUES ('Давид Алексеевич Клитин', 2);
INSERT INTO students (full_name, group_id) VALUES ('Деменьтий Сергеевич Остапов', 2);
INSERT INTO students (full_name, group_id) VALUES ('Сергей Николаевич Пегов', 2);
INSERT INTO students (full_name, group_id) VALUES ('Екатерина Сергеевна Бубнова', 2);

INSERT INTO students (full_name, group_id) VALUES ('Валерий Леонидович Антонов', 3);
INSERT INTO students (full_name, group_id) VALUES ('Денис Данилович Косыгин', 3);
INSERT INTO students (full_name, group_id) VALUES ('Елена Артёмовна Жигина', 3);
INSERT INTO students (full_name, group_id) VALUES ('Анастасия Олеговна Бугина', 3);
INSERT INTO students (full_name, group_id) VALUES ('Николай Дмитриевич Чугин', 3);
INSERT INTO students (full_name, group_id) VALUES ('Антон Александрович Серов', 3);

INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', '2023-10-21 09:00:00', 1, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', '2023-10-21 10:40:00', 2, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', '2023-10-23 14:20:00', 3, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', '2023-10-23 16:20:00', 4, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', '2023-10-25 12:40:00', 5, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', '2023-10-26 10:40:00', 6, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', '2023-10-26 12:40:00', 7, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', '2023-10-26 14:20:00', 8, 3, 1);

INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 1, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 2, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 3, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 4, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 5, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 6, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 7, 3);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 8, 3);

INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (2, 1, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (2, 2, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (2, 3, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (2, 4, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (2, 5, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (2, 6, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (2, 7, 3);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (2, 8, 3);

INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (3, 1, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (3, 2, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (3, 3, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (3, 4, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (3, 5, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (3, 6, 2);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (3, 7, 3);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (3, 8, 3);

INSERT INTO attendances (stud_id, sched_id) VALUES (1, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (1, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (1, 3);
INSERT INTO attendances (stud_id, sched_id) VALUES (1, 4);
INSERT INTO attendances (stud_id, sched_id) VALUES (1, 5);
INSERT INTO attendances (stud_id, sched_id) VALUES (1, 6);
INSERT INTO attendances (stud_id, sched_id) VALUES (1, 7);
INSERT INTO attendances (stud_id, sched_id) VALUES (1, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (2, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (2, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (2, 5);
INSERT INTO attendances (stud_id, sched_id) VALUES (2, 6);
INSERT INTO attendances (stud_id, sched_id) VALUES (2, 7);
INSERT INTO attendances (stud_id, sched_id) VALUES (2, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (3, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (3, 6);
INSERT INTO attendances (stud_id, sched_id) VALUES (3, 7);
INSERT INTO attendances (stud_id, sched_id) VALUES (3, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (4, 1);

INSERT INTO attendances (stud_id, sched_id) VALUES (5, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (5, 2);

INSERT INTO attendances (stud_id, sched_id) VALUES (6, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (6, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (6, 3);
INSERT INTO attendances (stud_id, sched_id) VALUES (6, 4);
INSERT INTO attendances (stud_id, sched_id) VALUES (6, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (7, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (7, 5);
INSERT INTO attendances (stud_id, sched_id) VALUES (7, 6);
INSERT INTO attendances (stud_id, sched_id) VALUES (7, 7);
INSERT INTO attendances (stud_id, sched_id) VALUES (7, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (8, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (8, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (8, 3);
INSERT INTO attendances (stud_id, sched_id) VALUES (8, 4);
INSERT INTO attendances (stud_id, sched_id) VALUES (8, 5);

INSERT INTO attendances (stud_id, sched_id) VALUES (9, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (9, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (9, 3);
INSERT INTO attendances (stud_id, sched_id) VALUES (9, 4);

INSERT INTO attendances (stud_id, sched_id) VALUES (10, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (10, 3);
INSERT INTO attendances (stud_id, sched_id) VALUES (10, 4);
INSERT INTO attendances (stud_id, sched_id) VALUES (10, 5);
INSERT INTO attendances (stud_id, sched_id) VALUES (10, 6);
INSERT INTO attendances (stud_id, sched_id) VALUES (10, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (11, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (11, 4);
INSERT INTO attendances (stud_id, sched_id) VALUES (11, 5);
INSERT INTO attendances (stud_id, sched_id) VALUES (11, 6);
INSERT INTO attendances (stud_id, sched_id) VALUES (11, 7);
INSERT INTO attendances (stud_id, sched_id) VALUES (11, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (12, 6);
INSERT INTO attendances (stud_id, sched_id) VALUES (12, 7);
INSERT INTO attendances (stud_id, sched_id) VALUES (12, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (13, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (13, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (13, 3);

INSERT INTO attendances (stud_id, sched_id) VALUES (14, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (15, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (15, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (15, 7);
INSERT INTO attendances (stud_id, sched_id) VALUES (15, 8);

INSERT INTO attendances (stud_id, sched_id) VALUES (16, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (16, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (16, 3);

INSERT INTO attendances (stud_id, sched_id) VALUES (17, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (17, 8);
					   