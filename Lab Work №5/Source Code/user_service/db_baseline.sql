USE userdb;

CREATE TABLE IF NOT EXISTS professor (
    id INT NOT NULL auto_increment,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    middlename VARCHAR(100),
    calendar_email VARCHAR(50) NOT NULL,
    calendar_api_key VARCHAR(100) NOT NULL,
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