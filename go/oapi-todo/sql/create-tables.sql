DROP TABLE IF EXISTS todo;

CREATE TABLE todo (
  id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  text text NOT NULL,
  done integer NOT NULL DEFAULT 0
);

INSERT INTO
  todo (text, done)
VALUES
  ('Do the Laundry', true),
  ('Clean the Kitchen', false),
  ('Mow the Lawn', false);
