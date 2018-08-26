CREATE SEQUENCE channel_ids;
CREATE TABLE channel (
  id INTEGER PRIMARY KEY DEFAULT NEXTVAL('channel_ids'),
  name CHAR(256),
  link CHAR(256),
  categories CHAR(256),
  subscribers INTEGER);