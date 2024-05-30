CREATE EXTENSION "uuid-ossp";

CREATE TABLE authors (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NULL,
    CONSTRAINT authors_pkey PRIMARY KEY (id)
);

CREATE TABLE books (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NULL,
    author_id UUID NOT NULL,
    rating FLOAT NOT NULL DEFAULT 0,
    published_at DATE NOT NULL,
    CONSTRAINT books_pkey PRIMARY KEY (id),
    CONSTRAINT books_author_id_fk FOREIGN KEY (author_id) REFERENCES authors (id)
);
