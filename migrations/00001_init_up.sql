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
    segment_id INTEGER NOT NULL REFERENCES segments (id) ON DELETE CASCADE,
    UNIQUE (user_id, segment_id)
);

CREATE TABLE IF NOT EXISTS segments_users_history (
    id SERIAL PRIMARY KEY UNIQUE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    segment_name TEXT NOT NULL,
    operation TEXT NOT NULL,
    operation_time DATE NOT NULL
);

create or replace function after_segment_deleted() returns trigger as $$
begin
    insert into segments_users_history
        (user_id, segment_name, operation, operation_time)
    values
        (user_id,  now());
    return old;
end;
$$ language plpgsql;

create trigger a_adr after delete on segments_users for each row execute procedure after_segment_deleted();