CREATE TABLE groups (
                        id SERIAL PRIMARY KEY,
                        group_name VARCHAR(20) NOT NULL,
                        faculty   VARCHAR(50) NOT NULL
);

INSERT INTO groups (group_name, faculty) VALUES
                                             ('CS-1', 'Computer Science'),
                                             ('EE-1', 'Electrical Engineering'),
                                             ('ME-1', 'Mechanical and Aerospace Engineering'),
                                             ('BIO-1', 'Biology');


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

INSERT INTO students (first_name, last_name, gender, birth_date, year_of_study, group_id)
VALUES
    ('Batyrkhan', 'Kamalov', 'M', '2003-05-10', 2, 1),
    ('Islam', 'Makhachev', 'M', '2002-11-21', 3, 2),
    ('Zhanna', 'Akhmetova', 'F', '2001-01-03', 4, 3),
    ('Daryn', 'Serik', 'F', '2001-06-18', 4, 3),
    ('Anna', 'Smirnova', 'F', '2003-03-12', 2, 1),
    ('Alexandr', 'Volkanovski', 'M', '2002-09-01', 3, 1);


CREATE TABLE courses (
                         id SERIAL PRIMARY KEY,
                         course_name VARCHAR(100) NOT NULL,
                         faculty     VARCHAR(50) NOT NULL,
                         credits     SMALLINT NOT NULL
);

INSERT INTO courses (course_name, faculty, credits)
VALUES
    ('Databases', 'Computer Science', 6),
    ('Algorithms', 'Computer Science', 8),
    ('Biology 101', 'Biology', 6),
    ('Linear Algebra', 'Engineering', 8);

CREATE TABLE student_courses (
                                 student_id INTEGER NOT NULL,
                                 course_id  INTEGER NOT NULL,

                                 PRIMARY KEY (student_id, course_id),

                                 FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
                                 FOREIGN KEY (course_id)  REFERENCES courses(id)  ON DELETE CASCADE
);

INSERT INTO student_courses VALUES
                                (1, 1),
                                (1, 2),
                                (2, 2),
                                (3, 3),
                                (4, 3),
                                (5, 1),
                                (6, 1),
                                (6, 2);

CREATE TABLE schedule (
                          id SERIAL PRIMARY KEY,
                          course_id INTEGER NOT NULL,
                          group_id INTEGER,
                          faculty VARCHAR(50),
                          day_of_week VARCHAR(10) NOT NULL,
                          start_time TIME NOT NULL,
                          end_time TIME NOT NULL,

                          FOREIGN KEY (course_id) REFERENCES courses(id),
                          FOREIGN KEY (group_id) REFERENCES groups(id)
);

INSERT INTO schedule (course_id, group_id, faculty, day_of_week, start_time, end_time)
VALUES
    (1, 1, 'Computer Science', 'Monday', '09:00', '10:30'),
    (2, 1, 'Computer Science', 'Monday', '10:45', '12:15'),
    (4, 2, 'Engineering', 'Tuesday', '09:00', '10:30'),
    (3, 4, 'Biology', 'Wednesday', '11:00', '12:30'),
    (1, 1, 'Computer Science', 'Thursday', '09:00', '10:30');

SELECT *
FROM students
WHERE gender = 'F'
ORDER BY birth_date DESC;

SELECT
    g.group_name,
    c.course_name,
    s.day_of_week,
    s.start_time,
    s.end_time
FROM schedule s
         JOIN courses c ON s.course_id = c.id
         JOIN groups g ON s.group_id = g.id
ORDER BY g.group_name, s.day_of_week, s.start_time;
