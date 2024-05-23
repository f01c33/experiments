CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name text not null,
    email text not null,
    password varchar(255) not null,
    code varchar(7),
    kind varchar(10) not null
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_name_user ON user(name);
CREATE UNIQUE INDEX IF NOT EXISTS idx_email_user ON user(email);

CREATE TABLE IF NOT EXISTS logins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user INTEGER NOT NULL,
    dt datetime NOT NULL,
    FOREIGN KEY (user)
        REFERENCES user (id)
);

CREATE TABLE IF NOT EXISTS posts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    body text not null,
    title text not null,
    dt datetime not null,
    category INTEGER NOT NULL,
    FOREIGN KEY (category)
        REFERENCES categories (id)
);

CREATE TABLE IF NOT EXISTS category(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_titulo_post ON posts(title);
