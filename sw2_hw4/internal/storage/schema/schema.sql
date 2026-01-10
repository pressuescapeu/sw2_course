CREATE TABLE faculty (
                         id SERIAL PRIMARY KEY,
                         faculty_name VARCHAR(40) NOT NULL
);

INSERT INTO faculty (faculty_name) VALUES
                                       ('Computer Science'),
                                       ('Electrical Engineering'),
                                       ('Mechanical and Aerospace Engineering'),
                                       ('Biology'),
                                       ('Mathematics');

CREATE TABLE groups (
                        id SERIAL PRIMARY KEY,
                        group_name VARCHAR(40) NOT NULL,
                        faculty_id INTEGER NOT NULL,

    FOREIGN KEY (faculty_id) REFERENCES faculty(id)
);

INSERT INTO groups (group_name, faculty_id) VALUES
                                             ('CS-1', 1),
                                             ('EE-1', 2),
                                             ('ME-1', 3),
                                             ('BIO-1', 4),
                                             ('MA-1', 5);


CREATE TABLE students (
                          id SERIAL PRIMARY KEY,
                          first_name VARCHAR(45) NOT NULL,
                          last_name  VARCHAR(45) NOT NULL,
                          gender     CHAR(1) CHECK (gender IN ('M', 'F')),
                          birth_date DATE NOT NULL,
                          year_of_study SMALLINT NOT NULL,
                          group_id INTEGER NOT NULL,

                          FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE courses (
                         id SERIAL PRIMARY KEY,
                         course_name VARCHAR(100) NOT NULL,
                         faculty_id     INTEGER NOT NULL,
                         credits     SMALLINT NOT NULL,

    FOREIGN KEY (faculty_id) REFERENCES faculty(id)
);

INSERT INTO courses (course_name, faculty_id, credits)
VALUES
    ('Databases', 1, 6),
    ('Circuits 101', 2, 8),
    ('Biology 101', 4, 6),
    ('Linear Algebra', 5, 8),
    ('Physics 101', 3, 6);

CREATE TABLE student_courses (
                                 student_id INTEGER NOT NULL,
                                 course_id  INTEGER NOT NULL,

                                 PRIMARY KEY (student_id, course_id),

                                 FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
                                 FOREIGN KEY (course_id)  REFERENCES courses(id)  ON DELETE CASCADE
);

CREATE TABLE schedule (
                          id SERIAL PRIMARY KEY,
                          course_id INTEGER NOT NULL,
                          group_id INTEGER,
                          date DATE NOT NULL,
                          start_time TIME NOT NULL,
                          end_time TIME NOT NULL,

                          FOREIGN KEY (course_id) REFERENCES courses(id),
                          FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE attendance (
                            id SERIAL PRIMARY KEY,
                            course_id INTEGER NOT NULL,
                            date DATE NOT NULL,
                            visited BOOLEAN NOT NULL,
                            student_id INTEGER NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(id),
    FOREIGN KEY (student_id) REFERENCES students(id)
);

-- from here
INSERT INTO students (first_name, last_name, gender, birth_date, year_of_study, group_id)
VALUES
    ('Lionel', 'Messi', 'M', '2003-06-24', 2, 1),
    ('Serena', 'Williams', 'F', '2002-09-26', 3, 2),
    ('Max', 'Verstappen', 'M', '2003-09-30', 2, 3),
    ('Simone', 'Biles', 'F', '2004-03-14', 1, 4),
    ('LeBron', 'James', 'M', '2002-12-30', 3, 5),
    ('Naomi', 'Osaka', 'F', '2003-10-16', 2, 1),
    ('Cristiano', 'Ronaldo', 'M', '2002-02-05', 3, 2),
    ('Katie', 'Ledecky', 'F', '2004-03-17', 1, 3),
    ('Usain', 'Bolt', 'M', '2001-08-21', 4, 4),
    ('Megan', 'Rapinoe', 'F', '2001-07-05', 4, 5);

INSERT INTO student_courses VALUES
                                (1, 1), (1, 4),
                                (2, 2), (2, 5),
                                (3, 5), (3, 1),
                                (4, 3), (4, 4),
                                (5, 4), (5, 1),
                                (6, 1), (6, 2),
                                (7, 2), (7, 5),
                                (8, 5), (8, 3),
                                (9, 3), (9, 4),
                                (10, 4), (10, 1);

INSERT INTO schedule (course_id, group_id, date, start_time, end_time)
VALUES
    -- CS-1 classes
    (1, 1, '2026-01-05', '09:00', '10:30'),
    (1, 1, '2026-01-12', '09:00', '10:30'),
    (1, 1, '2026-01-19', '09:00', '10:30'),
    (1, 1, '2026-01-26', '09:00', '10:30'),
    (1, 1, '2026-02-02', '09:00', '10:30'),

    -- EE-1 classes
    (2, 2, '2026-01-06', '10:45', '12:15'),
    (2, 2, '2026-01-13', '10:45', '12:15'),
    (2, 2, '2026-01-20', '10:45', '12:15'),
    (2, 2, '2026-01-27', '10:45', '12:15'),
    (2, 2, '2026-02-03', '10:45', '12:15'),

    -- ME-1 classes
    (5, 3, '2026-01-07', '14:00', '15:30'),
    (5, 3, '2026-01-14', '14:00', '15:30'),
    (5, 3, '2026-01-21', '14:00', '15:30'),
    (5, 3, '2026-01-28', '14:00', '15:30'),
    (5, 3, '2026-02-04', '14:00', '15:30'),

    -- BIO-1 classes
    (3, 4, '2026-01-08', '11:00', '12:30'),
    (3, 4, '2026-01-15', '11:00', '12:30'),
    (3, 4, '2026-01-22', '11:00', '12:30'),
    (3, 4, '2026-01-29', '11:00', '12:30'),
    (3, 4, '2026-02-05', '11:00', '12:30'),

    -- MA-1 classes
    (4, 5, '2026-01-09', '09:00', '10:30'),
    (4, 5, '2026-01-16', '09:00', '10:30'),
    (4, 5, '2026-01-23', '09:00', '10:30'),
    (4, 5, '2026-01-30', '09:00', '10:30'),
    (4, 5, '2026-02-06', '09:00', '10:30');

-- to here - this part is AI-generated

SELECT *
FROM students
WHERE gender = 'F'
ORDER BY birth_date DESC;

SELECT
    g.group_name,
    c.course_name,
    s.date,
    s.start_time,
    s.end_time
FROM schedule s
         JOIN courses c ON s.course_id = c.id
         JOIN groups g ON s.group_id = g.id
ORDER BY g.group_name, s.date, s.start_time;
