CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO users (username, password) VALUES
('alice', 'password123', NOW(), NOW()),
('bob', 'securepass456', NOW(), NOW()),
('charlie', 'qwerty789', NOW(), NOW()),
('david', 'letmein321', NOW(), NOW()),
('eve', 'passw0rd', NOW(), NOW());
