CREATE TABLE IF NOT EXISTS departments (
    department_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT,
    leader TEXT,
    website TEXT,
    faculty TEXT,
    base_organizations TEXT
);

CREATE TABLE IF NOT EXISTS subjects (
    subject_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS teachers (
    teacher_id SERIAL PRIMARY KEY,
    name TEXT,
    birth_date DATE,
    alma_mater TEXT, 
    graduation_year DATE,
    degree TEXT,
    knowledge_score DECIMAL(3, 2),
    teaching_skill_score DECIMAL(3, 2),
    communication_score DECIMAL(3, 2),
    leniency_score DECIMAL(3, 2),
    overall_score DECIMAL(3, 2),
    knowledge_rating_num SMALLINT,
    teaching_skill_rating_num SMALLINT,
    communication_rating_num SMALLINT,
    leniency_rating_num SMALLINT,
    overall_rating_num SMALLINT
);

CREATE TABLE IF NOT EXISTS subject_teacher (
    subject_id INTEGER REFERENCES subjects(subject_id),
    teacher_id INTEGER REFERENCES teachers(teacher_id),
    PRIMARY KEY (subject_id, teacher_id)
);

CREATE TABLE IF NOT EXISTS department_teacher (
    department_id INTEGER REFERENCES departments(department_id),
    teacher_id INTEGER REFERENCES teachers(teacher_id),
    PRIMARY KEY (department_id, teacher_id)
);

CREATE TABLE IF NOT EXISTS department_subject (
    department_id INTEGER REFERENCES departments(department_id),
    subject_id INTEGER REFERENCES subjects(subject_id),
    PRIMARY KEY (department_id, subject_id)
);

TRUNCATE TABLE department_subject CASCADE;
TRUNCATE TABLE department_teacher CASCADE;
TRUNCATE TABLE subject_teacher CASCADE;
TRUNCATE TABLE teachers CASCADE;
TRUNCATE TABLE subjects CASCADE;
TRUNCATE TABLE departments CASCADE;
    
ALTER SEQUENCE departments_department_id_seq RESTART WITH 1;
ALTER SEQUENCE subjects_subject_id_seq RESTART WITH 1;
ALTER SEQUENCE teachers_teacher_id_seq RESTART WITH 1;


# Введите ваш путь вместо {YOUR_PATH}

\COPY tickets FROM ‘{YOUR_PATH}/tg_bot/departments.csv’ DELIMITER ‘,’ CSV HEADER;
\COPY tickets FROM ‘{YOUR_PATH}/tg_bot/department_subject.csv’ DELIMITER ‘,’ CSV HEADER;
\COPY tickets FROM ‘{YOUR_PATH}/tg_bot/department_teacher.csv’ DELIMITER ‘,’ CSV HEADER;
\COPY tickets FROM ‘{YOUR_PATH}/tg_bot/subject_teacher.csv’ DELIMITER ‘,’ CSV HEADER;
\COPY tickets FROM ‘{YOUR_PATH}/tg_bot/subjects.csv’ DELIMITER ‘,’ CSV HEADER;
\COPY tickets FROM ‘{YOUR_PATH}/tg_bot/teachers.csv’ DELIMITER ‘,’ CSV HEADER;