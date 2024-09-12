CREATE TABLE IF NOT EXISTS comment (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    user        INTEGER,
    post        INTEGER,
    text        VARCHAR(512),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user) REFERENCES user(id),
    FOREIGN KEY (post) REFERENCES post(id)
)