CREATE TABLE IF NOT EXISTS invite (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    code        VARCHAR(255),
    created_by  INTEGER,
    used_by     INTEGER,

    FOREIGN KEY (created_by) REFERENCES user(id),
    FOREIGN KEY (used_by) REFERENCES user(id)
)