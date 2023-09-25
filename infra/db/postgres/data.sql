CREATE TABLE users (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    nik VARCHAR(50) UNIQUE,
    phone_number VARCHAR(50) UNIQUE,
    pin_number VARCHAR(10),
    account_number VARCHAR(50) UNIQUE,
    balance BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE INDEX idx_user_id ON users (id);
CREATE INDEX idx_user_data_id ON users (account_number);

CREATE TABLE mutation_data (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    code_transaction VARCHAR(255),
    account_number VARCHAR(50),
    transaction VARCHAR(50),
    transaction_at TIMESTAMP,
    
    CONSTRAINT fk_user_data_account_number FOREIGN KEY (account_number) REFERENCES users (account_number)
);

CREATE INDEX idx_mutation_data_id ON mutation_data (id);
CREATE INDEX idx_mutation_data_account_number ON mutation_data (account_number);
