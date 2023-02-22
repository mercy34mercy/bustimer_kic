CREATE TABLE busstop_url (
  id   INTEGER PRIMARY KEY,
  busstop VARCHAR(255) NOT NUll,
  busname VARCHAR(255) NOT NUll,
  destination VARCHAR(255) NOT NUll,
  url VARCHAR(255) UNIQUE NOT NUll
);