CREATE TABLE IF NOT EXISTS user (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        VARCHAR(255),
    password    VARCHAR(512)
)