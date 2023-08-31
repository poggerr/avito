CREATE TABLE IF NOT EXISTS segments (
    id UUID UNIQUE PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS users (
    id UUID UNIQUE PRIMARY KEY,
    username VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS user_segments (
    id INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    segment_slug VARCHAR(255) REFERENCES segments(slug) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS segment_history (
    user_id UUID REFERENCES users(id),
    segment_slug VARCHAR(255) REFERENCES segments(slug),
    operation VARCHAR(255),
    datetime VARCHAR(255)
);

INSERT INTO users (id, username) VALUES ('2376e110-e40d-41d0-85ba-22db804c4f51', 'Test1');