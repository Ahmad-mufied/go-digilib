DROP DATABASE IF EXISTS go_digilib_dev;
CREATE DATABASE go_digilib_dev;


/*/ DROP ALL TABLES and ENUMS */
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS borrow_prices;
DROP TABLE IF EXISTS borrows;
DROP TABLE IF EXISTS book_physical_details;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS deposits;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS PaymentStatusEnum;
DROP TYPE IF EXISTS UserRoleEnum;
DROP TYPE IF EXISTS DurationType;
DROP TYPE IF EXISTS BorrowStatusEnum;
DROP TYPE IF EXISTS UsersStatusEnum;
DROP TYPE IF EXISTS BookStatusEnum;

-- Create Enums
CREATE TYPE UsersStatusEnum AS ENUM ('active', 'inactive', 'banned');
CREATE TYPE BorrowStatusEnum AS ENUM ('returned', 'success', 'pending', 'cancel');
CREATE TYPE DurationType AS ENUM ('daily', 'weekly', 'monthly');
CREATE TYPE UserRoleEnum AS ENUM ('reader', 'admin');
CREATE TYPE PaymentStatusEnum AS ENUM ('pending', 'confirmed', 'refunded');
CREATE TYPE BookStatusEnum AS ENUM ('available', 'not available');

-- Create Tables
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    full_name  VARCHAR(100)    NOT NULL,
    username   VARCHAR(100)    NOT NULL,
    email      VARCHAR(100)    NOT NULL,
    password   VARCHAR(100)    NOT NULL,
    status     UsersStatusEnum NOT NULL DEFAULT 'inactive',
    role       UserRoleEnum    NOT NULL DEFAULT 'reader',
    book_count INT             NOT NULL DEFAULT 0,
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE wallets
(
    id         SERIAL PRIMARY KEY,
    user_id    INT            NOT NULL,
    balance    DECIMAL(10, 2) NOT NULL CHECK ( balance >= 0 ),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE deposits
(
    id          SERIAL PRIMARY KEY,
    amount      DECIMAL(10, 2)    NOT NULL CHECK ( amount >= 0 ),
    wallet_id   INT               NOT NULL,
    invoice_url TEXT              NOT NULL,
    status      PaymentStatusEnum NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMP                  DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE books
(
    id             SERIAL PRIMARY KEY,
    category_id    INT,
    isbn           VARCHAR(50)    NOT NULL,
    sku            VARCHAR(50)    NOT NULL,
    author         JSON           NOT NULL,
    title          VARCHAR(120)   NOT NULL,
    image          VARCHAR(255)   NOT NULL,
    pages          SMALLINT       NOT NULL DEFAULT 0,
    language       VARCHAR(20)    NOT NULL,
    description    TEXT           NOT NULL,
    stock          SMALLINT       NOT NULL DEFAULT 0,
    status         BookStatusEnum NOT NULL DEFAULT 'not available',
    borrowed_count INT            NOT NULL DEFAULT 0,
    published_at   TIMESTAMP               DEFAULT NULL,
    base_price     DECIMAL(10, 2) NOT NULL CHECK ( base_price >= 0 ),
    created_at     TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE book_physical_details
(
    id         SERIAL PRIMARY KEY,
    book_id    INT      NOT NULL,
    weight     FLOAT    NOT NULL DEFAULT 0,
    height     SMALLINT NOT NULL DEFAULT 0,
    width      SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP         DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE borrows
(
    id          SERIAL PRIMARY KEY,
    user_id     INT              NOT NULL,
    book_id     INT              NOT NULL,
    status      BorrowStatusEnum NOT NULL DEFAULT 'pending',
    start_date  TIMESTAMP        NOT NULL,
    end_date    TIMESTAMP        NOT NULL,
    total_price DECIMAL(10, 2)   NOT NULL CHECK ( total_price >= 0 ),
    returned_at TIMESTAMP                 DEFAULT NULL,
    created_at  TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                 DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE borrow_prices
(
    id               SERIAL PRIMARY KEY,
    book_id          INT          NOT NULL,
    duration_type    DurationType NOT NULL,
    price_multiplier DECIMAL      NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payments
(
    id           SERIAL PRIMARY KEY,
    borrow_id    INT,
    amount       DECIMAL(10, 2)    NOT NULL CHECK ( amount >= 0 ),
    payment_date TIMESTAMP         NOT NULL,
    status       PaymentStatusEnum NOT NULL DEFAULT 'pending',
    created_at   TIMESTAMP                  DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP                  DEFAULT CURRENT_TIMESTAMP
);

-- Add Foreign Key Constraints
ALTER TABLE wallets
    ADD CONSTRAINT fk_wallets_users
        FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE deposits
    ADD CONSTRAINT fk_deposits_wallets
        FOREIGN KEY (wallet_id) REFERENCES wallets (id);

ALTER TABLE books
    ADD CONSTRAINT fk_books_categories
        FOREIGN KEY (category_id) REFERENCES categories (id);

ALTER TABLE book_physical_details
    ADD CONSTRAINT fk_book_physical_details_books
        FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE;

ALTER TABLE borrows
    ADD CONSTRAINT fk_borrows_users
        FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE borrows
    ADD CONSTRAINT fk_borrows_books
        FOREIGN KEY (book_id) REFERENCES books (id);

ALTER TABLE borrow_prices
    ADD CONSTRAINT fk_borrow_prices_books
        FOREIGN KEY (book_id) REFERENCES books (id);

ALTER TABLE payments
    ADD CONSTRAINT fk_payments_borrows
        FOREIGN KEY (borrow_id) REFERENCES borrows (id);
