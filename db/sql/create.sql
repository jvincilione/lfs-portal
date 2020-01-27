CREATE DATABASE IF NOT EXISTS labonte;

USE labonte;

CREATE TABLE IF NOT EXISTS customers (
  id BIGINT AUTO_INCREMENT KEY,
  email VARCHAR(255) NOT NULL,
  full_name VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  address2 VARCHAR(255),
  city VARCHAR(75) NOT NULL,
  state VARCHAR(13) NOT NULL,
  postal_code VARCHAR(255) NOT NULL,
  phone_number VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS jobs (
  id BIGINT AUTO_INCREMENT KEY,
  customer_id BIGINT,
  full_name VARCHAR(255),
  address VARCHAR(255),
  address2 VARCHAR(255),
  city VARCHAR(75),
  state VARCHAR(13),
  postal_code VARCHAR(255),
  order_number VARCHAR(255),
  instructions VARCHAR(255),
  submitted_date DATETIME,
  scheduled_date DATETIME,
  status VARCHAR(255),
  parts_cost DECIMAL(13, 2),
  labor_cost DECIMAL(13, 2),
  FOREIGN KEY fk_customer(customer_id) REFERENCES customers(id)
);

CREATE TABLE IF NOT EXISTS phone_numbers (
  id BIGINT AUTO_INCREMENT KEY,
  job_id BIGINT,
  phone_number VARCHAR(255),
  type INT,
  FOREIGN KEY fk_job(job_id) REFERENCES jobs(id)
);

CREATE TABLE IF NOT EXISTS notes (
  id BIGINT AUTO_INCREMENT KEY,
  job_id BIGINT,
  notes TEXT,
  type INT,
  created_date DATETIME,
  updated_date DATETIME,
  FOREIGN KEY fk_job(job_id) REFERENCES jobs(id)
);

CREATE TABLE IF NOT EXISTS resources (
  id BIGINT AUTO_INCREMENT KEY,
  job_id BIGINT,
  url TEXT,
  type INT,
  FOREIGN KEY fk_job(job_id) REFERENCES jobs(id)
);

CREATE TABLE IF NOT EXISTS users (
  id BIGINT AUTO_INCREMENT KEY,
  customer_id BIGINT,
  email VARCHAR(255) NOT NULL,
  user_type INT NOT NULL,
  password VARCHAR(64) NOT NULL,
  salt VARCHAR(255) NOT NULL,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  FOREIGN KEY fk_customer(customer_id) REFERENCES customers(id)
);

CREATE TABLE IF NOT EXISTS messages (
  id BIGINT AUTO_INCREMENT KEY,
  user_id BIGINT NOT NULL,
  recipient_id BIGINT NOT NULL,
  message TEXT NOT NULL,
  status INT,
  FOREIGN KEY fk_user(user_id) REFERENCES users(id),
  FOREIGN KEY fk_recipient(recipient_id) REFERENCES users(id)
);
