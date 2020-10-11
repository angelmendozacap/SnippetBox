CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password CHAR(60) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uq_email UNIQUE (email);

INSERT INTO users (name, email, password, created_at) VALUES (
    'Alice Jones',
    'alice@example.com',
    '$2a$12$Z6l.CHyOXGnsHo5rK1zi8eLHa.TtSYdpWIhSWwiMkX0ffKzOtxSEm',
    '2018-12-23 17:25:22'
);

CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    user_id integer not null,
    content TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    expires_at DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created_at ON snippets(created_at);

ALTER TABLE snippets ADD CONSTRAINT snippets_fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE;
