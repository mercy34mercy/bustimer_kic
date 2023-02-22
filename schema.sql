CREATE TABLE busstop_url (
  id   INTEGER PRIMARY KEY,
  busstop VARCHAR(255),
  busname VARCHAR(255),
  destination VARCHAR(255),
  url VARCHAR(255) UNIQUE
);