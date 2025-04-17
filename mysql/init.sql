DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO users (username, password, created_at, updated_at) VALUES
('alice', 'password123', NOW(), NOW()),
('bob', 'securepass456', NOW(), NOW()),
('charlie', 'qwerty789', NOW(), NOW()),
('david', 'letmein321', NOW(), NOW()),
('eve', 'passw0rd', NOW(), NOW());