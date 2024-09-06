CREATE TABLE IF NOT EXISTS session (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    user        INTEGER,
    created_at  TIMESTAMP,
    expires_at  TIMESTAMP,

    FOREIGN KEY (user) REFERENCES user(id)
)