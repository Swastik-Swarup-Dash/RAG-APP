CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    embedding VECTOR(768) NOT NULL -- Adjust dimension based on Gemini model used.
);

CREATE INDEX ON documents USING ivfflat (embedding vector_l2_ops);
