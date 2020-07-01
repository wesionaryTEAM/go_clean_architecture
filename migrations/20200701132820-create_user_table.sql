
-- +migrate Up
CREATE TABLE contacts
(
  id  INT(20) NOT NULL AUTO_INCREMENT,
  contact_number VARCHAR(255),
  CONSTRAINT contacts_pk PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE contacts;