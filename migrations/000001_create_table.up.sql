create type  genders as enum('mele','femele');

CREATE TABLE IF NOT EXISTS allusers(
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(30) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(50) NOT NULL,
    gender genders NOT NULL,
    password varchar(30) not null,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at BIGINT DEFAULT 0 NOT NULL
);

