CREATE TABLE IF NOT EXISTS session (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    user        INTEGER,
    uuid        VARCHAR(36),
    created_at  TIMESTAMP,
    expires_at  TIMESTAMP,

    FOREIGN KEY (user) REFERENCES user(id)
)