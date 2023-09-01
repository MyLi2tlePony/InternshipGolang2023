CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    user_id INTEGER UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS segments(
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS segments_users(
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    segment_id INTEGER NOT NULL REFERENCES segments (id),
    UNIQUE (user_id, segment_id),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS segments_users_history (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    segment_name TEXT NOT NULL,
    operation TEXT NOT NULL,
    operation_time TIMESTAMP NOT NULL
);

CREATE OR REPLACE FUNCTION user_segment_deleted() RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO segments_users_history
    (user_id, segment_name, operation, operation_time)
    VALUES
        (old.user_id,  (SELECT name FROM segments WHERE id = old.segment_id), 'удаление', now());
    RETURN old;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER after_segment_deleted
    AFTER DELETE ON segments_users
    FOR EACH ROW
EXECUTE PROCEDURE user_segment_deleted();

CREATE OR REPLACE FUNCTION user_segment_inserted() RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO segments_users_history
    (user_id, segment_name, operation, operation_time)
    VALUES
        (new.user_id,  (SELECT name FROM segments WHERE id = new.segment_id), 'добавление', now());
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER after_segment_inserted
    AFTER INSERT ON segments_users
    FOR EACH ROW
EXECUTE PROCEDURE user_segment_inserted();

CREATE OR REPLACE FUNCTION segment_deleted() RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM segments_users WHERE segments_users.segment_id = old.id;
    RETURN old;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER before_segment_deleted
    BEFORE DELETE ON segments
    FOR EACH ROW
EXECUTE PROCEDURE segment_deleted();