DROP TABLE IF EXISTS ADMINS;
DROP TABLE IF EXISTS TERMINALS;

CREATE TABLE ADMINS (
    admin_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('superadmin', 'operator terminal')),
    last_sync TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE TERMINALS (
    terminal_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

INSERT INTO ADMINS (username, password_hash, role) VALUES
('superadmin', '$2a$10$hPOLW/9Qz9oHFSKp8manruSxZC3ZXABXfwCm4oYHWIoLHXV38vqFG', 'superadmin');

SELECT 'Skema database E-Ticketing berhasil dibuat.' as "Status";