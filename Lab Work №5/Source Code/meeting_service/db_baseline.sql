USE meeting;

CREATE TABLE IF NOT EXISTS meeting (
    id INT NOT NULL auto_increment,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(100) NOT NULL,
    meeting_time DATETIME NOT NULL, 
    student_participant_id INT NOT NULL,
    is_online BOOLEAN NOT NULL,
    professor_id INT NOT NULL,
    PRIMARY KEY (id)
);