-- +goose Up
CREATE TABLE company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE TABLE package (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    company_id INTEGER REFERENCES company(id) ON DELETE CASCADE NOT NULL,
    price DECIMAL NULL,
    terms TEXT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE TABLE service (
    id SERIAL PRIMARY KEY,
    company_id INTEGER REFERENCES company(id) ON DELETE CASCADE NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE TABLE policyholder (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    telnumber VARCHAR(50) NOT NULL,
    code INTEGER NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE TABLE subscription (
    id SERIAL PRIMARY KEY,
    policyholder_id INTEGER REFERENCES policyholder(id) ON DELETE CASCADE NOT NULL,
    package_id INTEGER REFERENCES package(id) ON DELETE CASCADE NOT NULL,
    credit DECIMAL NULL,
    expires_on TIMESTAMP NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE TABLE Credit (
    id SERIAL PRIMARY KEY,
    subscription_id INTEGER REFERENCES subscription(id) ON DELETE CASCADE NOT NULL,
    amount DECIMAL NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE TABLE claim (
    id SERIAL PRIMARY KEY,
    policyholder_id INTEGER REFERENCES policyholder(id) ON DELETE CASCADE NOT NULL,
    reason VARCHAR(256) NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);