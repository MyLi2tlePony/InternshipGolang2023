DROP TABLE IF EXISTS users, segments, segments_users, segments_users_history;

DROP FUNCTION IF EXISTS user_segment_inserted() CASCADE;
DROP FUNCTION IF EXISTS user_segment_deleted() CASCADE;
DROP FUNCTION IF EXISTS segment_deleted() CASCADE;