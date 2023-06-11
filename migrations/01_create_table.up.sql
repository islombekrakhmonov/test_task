CREATE extension IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user" (
    user_id UUID  NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL, 
    token VARCHAR(255),
    refresh_token VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "phone" (
    phone_id UUID  NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    phone_number VARCHAR(12) NOT NULL UNIQUE,
    user_id UUID,
    description VARCHAR(255) NOT NULL,
    is_fax BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES "user" (user_id) ON DELETE CASCADE ON UPDATE CASCADE
);
