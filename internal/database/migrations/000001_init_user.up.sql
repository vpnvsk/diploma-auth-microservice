CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "role" AS ENUM ('admin', 'user');

CREATE TYPE auth_method AS ENUM ('email', 'oauth');

CREATE TABLE "user" (
                        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                        username VARCHAR(255) NOT NULL,
                        email VARCHAR(255) NOT NULL UNIQUE,
                        password_hash BYTEA,
                        role "role" NOT NULL DEFAULT 'user',
                        is_active BOOLEAN NOT NULL DEFAULT TRUE,
                        qa_thread_id VARCHAR(255),
                        researcher_thread_id VARCHAR(255),
                        vector_store_id VARCHAR(255),
                        newsletter BOOLEAN NOT NULL DEFAULT FALSE,
                        photo_url VARCHAR(255),
                        auth_method auth_method NOT NULL DEFAULT 'email',
                        refresh_token BYTEA,
                        created_at DATE NOT NULL DEFAULT CURRENT_DATE,
                        main_set UUID
);

CREATE INDEX idx_user_email ON "user" (email);
