CREATE TYPE CourseStatus AS ENUM('active', 'inactive');

CREATE TABLE Major (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(250) NOT NULL
);

CREATE TABLE UserAuthentication (
    id SERIAL PRIMARY KEY,
    username VARCHAR(60) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    last_login TIMESTAMP,
    last_password_reset TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE UserSession (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    username VARCHAR(60) NOT NULL, -- Not a 100% necessity becasue we have user_id. However, we use this to reduce db calls to refresh token
    refresh_token TEXT NOT NULL,
    refresh_token_id TEXT NOT NULL,
    issued_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked BOOLEAN DEFAULT FALSE,

    FOREIGN KEY (user_id) REFERENCES UserAuthentication(id)
);

CREATE TABLE Student (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    auth_id INT,
    major_id INT,
    phone_number VARCHAR(20) NOT NULL,
    email VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(major_id) REFERENCES Major(id),
    FOREIGN KEY(auth_id) REFERENCES UserAuthentication(id)
);

CREATE TABLE Teacher (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    auth_id INT,
    phone_number VARCHAR(20) NOT NULL,
    email VARCHAR(50) NOT NULL,
    office VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(auth_id) REFERENCES UserAuthentication(id)
);

CREATE TABLE Administrator (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    auth_id INT,
    phone_number VARCHAR(20) NOT NULL,
    office VARCHAR(60) NOT NULL,
    email VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(auth_id) REFERENCES UserAuthentication(id)
);


CREATE TABLE Course (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(255) NOT NULL,
    teacher_id INT,
    max_enrollment INT NOT NULL,
    num_credits INT NOT NULL,
    status CourseStatus NOT NULL,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(teacher_id) REFERENCES Teacher(id),
    CHECK ((status = 'inactive' AND max_enrollment = 0) OR status = 'active')
);

CREATE TABLE Course_Major (
    id SERIAL PRIMARY KEY,
    major_id INT,
    course_id INT,

    FOREIGN KEY(major_id) REFERENCES Major(id),
    FOREIGN KEY(course_id) REFERENCES Course(id)
);

CREATE TABLE Registered (
    id SERIAL PRIMARY KEY,
    course_id INT,
    student_id INT,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    unregistered_date TIMESTAMP,
    is_unregistered BOOLEAN,

    FOREIGN KEY(course_id) REFERENCES Course(id),
    FOREIGN KEY(student_id) REFERENCES Student(id),
    CHECK ((is_unregistered = FALSE AND unregistered_date IS NULL) OR (is_unregistered = TRUE AND unregistered_date IS NOT NULL))
);

