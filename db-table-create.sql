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