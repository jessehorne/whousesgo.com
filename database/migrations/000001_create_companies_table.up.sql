CREATE TABLE IF NOT EXISTS companies (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    name VARCHAR(255),
    image VARCHAR(255),
    slug VARCHAR(255),
    description TEXT,
    location VARCHAR(255),
    website VARCHAR(255),
    github VARCHAR(255),
    linkedin VARCHAR(255)
);