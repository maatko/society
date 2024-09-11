CREATE TABLE IF NOT EXISTS post (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    user        INTEGER,
    uuid        VARCHAR(36),
    cover       VARCHAR(1024),
    about       VARCHAR(512),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user) REFERENCES user(id)
)