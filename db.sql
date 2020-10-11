DROP DATABASE IF EXISTS snippetbox;

CREATE DATABASE IF NOT EXISTS snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE snippetbox;

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password CHAR(60) NOT NULL,
    created_at DATETIME NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uq_email UNIQUE (email);

INSERT INTO users (name, email, password, created_at) VALUES (
    'Alice Jones',
    'alice@example.com',
    '$2a$12$Z6l.CHyOXGnsHo5rK1zi8eLHa.TtSYdpWIhSWwiMkX0ffKzOtxSEm',
    '2018-12-23 17:25:22'
);

-- Create a `snippets` table.
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    expires_at DATETIME NOT NULL
);

ALTER TABLE snippets ADD CONSTRAINT snippets_fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE;

CREATE INDEX idx_snippets_created ON snippets(created_at);

INSERT INTO snippets (title, content, expires_at, user_id) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY),
    1
);

INSERT INTO snippets (title, content, expires_at, user_id) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY),
    1
);

INSERT INTO snippets (title, content, expires_at, user_id) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY),
    1
);
