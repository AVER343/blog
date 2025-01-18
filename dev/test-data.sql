
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE
    IF NOT EXISTS users (
        id VARCHAR(127) PRIMARY KEY DEFAULT gen_random_uuid (),
        username VARCHAR(64) NOT NULL UNIQUE,
        password VARCHAR(60) NOT NULL,
        email citext NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
CREATE TABLE
    IF NOT EXISTS posts (
        id BIGSERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        user_id VARCHAR(127) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );


-- Insert sample data into the users table
INSERT INTO users (id, username, password, email, created_at, updated_at)
VALUES 
    (gen_random_uuid(), 'john_doe', 'hashed_password_1', 'john.doe@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'jane_smith', 'hashed_password_2', 'jane.smith@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'charlie_brown', 'hashed_password_3', 'charlie.brown@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert sample data into the posts table
INSERT INTO posts (title, content, user_id, created_at, updated_at)
VALUES 
    ('Post 1 Title', 'Content for post 1', (SELECT id FROM users WHERE username = 'john_doe'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Post 2 Title', 'Content for post 2', (SELECT id FROM users WHERE username = 'jane_smith'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Post 3 Title', 'Content for post 3', (SELECT id FROM users WHERE username = 'charlie_brown'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Post 4 Title', 'Content for post 4', (SELECT id FROM users WHERE username = 'john_doe'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);