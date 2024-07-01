DROP TABLE IF EXISTS todo;
CREATE TABLE todo (
  id      INT AUTO_INCREMENT NOT NULL,
  text    VARCHAR(255) NOT NULL,
  done    BOOLEAN NOT NULL DEFAULT false,
  PRIMARY KEY (id)
);

INSERT INTO todo
  (text, done)
VALUES
  ('Do the Laundry', true),
  ('Clean the Kitchen', false),
  ('Mow the Lawn', false);
