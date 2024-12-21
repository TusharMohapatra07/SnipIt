-- create snippets table
CREATE TABLE snippets(
	id SERIAL PRIMARY KEY NOT NULL,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	created TIMESTAMP WITH TIME ZONE NOT NULL,
	expires TIMESTAMP WITH TIME ZONE NOT NULL
);

-- add an index on created column of snippets table
CREATE INDEX ON snippets(created);

-- OPTIONAL: fill with dummy data

-- INSERT INTO snippets (title, content, created, expires) VALUES (
-- 'An old silent pond',
-- 'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
-- CURRENT_TIMESTAMP,
-- DATE_ADD(CURRENT_TIMESTAMP, '1 year'::interval)
-- ), (
-- 'Over the wintry forest',
-- 'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
-- CURRENT_TIMESTAMP,
-- DATE_ADD(CURRENT_TIMESTAMP, '1 year'::interval)
-- ), (
-- 'First autumn morning',
-- 'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
-- CURRENT_TIMESTAMP,
-- DATE_ADD(CURRENT_TIMESTAMP, '1 week'::interval)
-- );