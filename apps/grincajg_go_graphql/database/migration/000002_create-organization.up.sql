CREATE TABLE organizations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    admin_user_id INT UNIQUE,
    latitude NUMERIC(10, 8),  
    longitude NUMERIC(11, 8),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
