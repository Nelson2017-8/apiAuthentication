CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(50) NOT NULL,
    address VARCHAR(255) NOT NULL,
    password VARCHAR(350) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    token VARCHAR(350) NOT NULL,
    session_active BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT users_email_key UNIQUE (email)
    );

ALTER TABLE IF EXISTS public.users
    OWNER TO postgres;
