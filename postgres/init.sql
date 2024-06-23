CREATE DATABASE "graphqlDb";

\connect "graphqlDb"

DROP TABLE IF EXISTS posts CASCADE;

CREATE TABLE posts (
    id VARCHAR(30) PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    comments_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS comments CASCADE;

CREATE TABLE comments (
    id VARCHAR(30) PRIMARY KEY,
    post_id VARCHAR(30) NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    parent_id VARCHAR(30) REFERENCES comments(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- CREATE INDEX IF NOT EXISTS idx_comments_post_id_created_at ON comments (post_id, created_at);