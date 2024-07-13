CREATE TABLE students (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    subject TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
COPY students(id, name, subject, body, created_at, updated_at)
FROM '/docker-entrypoint-initdb.d/message_data.csv'
DELIMITER ','
CSV HEADER;