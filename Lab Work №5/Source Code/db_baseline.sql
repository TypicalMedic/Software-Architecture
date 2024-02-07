USE projectmanagement;

CREATE TABLE IF NOT EXISTS professor (
    id INT NOT NULL auto_increment,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    middlename VARCHAR(100),
    calendar_email VARCHAR(50) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS student (
    id INT NOT NULL auto_increment,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    middlename VARCHAR(100),
    cource INT NOT NULL,
    project_theme VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS meeting (
    id INT NOT NULL auto_increment,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(100) NOT NULL,
    meeting_time DATETIME NOT NULL, 
    student_participant_id INT NOT NULL,
    is_online BOOLEAN NOT NULL,
    professor_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (student_participant_id) REFERENCES student(id) ON DELETE CASCADE,
    FOREIGN KEY (professor_id) REFERENCES professor(id) ON DELETE CASCADE
);