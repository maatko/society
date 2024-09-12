CREATE TABLE IF NOT EXISTS like (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    user        INTEGER,
    post        INTEGER,

    FOREIGN KEY (user) REFERENCES user(id),
    FOREIGN KEY (post) REFERENCES post(id)
)