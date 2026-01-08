CREATE TABLE groups (
                        id SERIAL PRIMARY KEY,
                        group_name VARCHAR(20) NOT NULL,
                        faculty   VARCHAR(50) NOT NULL
);

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
                         faculty     VARCHAR(50) NOT NULL,
                         credits     SMALLINT NOT NULL
);

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
                          faculty VARCHAR(50),
                          day_of_week VARCHAR(10) NOT NULL,
                          start_time TIME NOT NULL,
                          end_time TIME NOT NULL,

                          FOREIGN KEY (course_id) REFERENCES courses(id),
                          FOREIGN KEY (group_id) REFERENCES groups(id)
);