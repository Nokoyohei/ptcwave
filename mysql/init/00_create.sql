CREATE DATABASE IF NOT EXISTS ptcwave;
USE ptcwave;
DROP TABLE IF EXISTS rush_cards;
CREATE TABLE IF NOT EXISTS rush_cards(
  name VARCHAR(50) NOT NULL,
  version VARCHAR(100) NOT NULL,
  rarity VARCHAR(10) NOT NULL,
  type VARCHAR(10) NOT NULL,
  price decimal NOT NULL,
  status VARCHAR(255) NOT NULL DEFAULT '-',
  fetch_date datetime NOT NULL DEFAULT current_timestamp,
  created_at datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY(name, version, rarity, fetch_date)
);

DROP TABLE IF EXISTS hareruya_cards;
CREATE TABLE IF NOT EXISTS hareruya_cards(
  name VARCHAR(50) NOT NULL,
  version VARCHAR(100) NOT NULL,
  rarity VARCHAR(50) NOT NULL,
  type VARCHAR(10) NOT NULL,
  price decimal NOT NULL,
  status VARCHAR(255) NOT NULL DEFAULT '-',
  fetch_date datetime NOT NULL DEFAULT current_timestamp,
  created_at datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY(name, version, rarity, status, fetch_date)
);