-----------------
-- Main Tables --
-----------------

-- Users table
CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    birthdate DATE NOT NULL,
    date_joined DATE NOT NULL
);

-- Groups table
CREATE TABLE IF NOT EXISTS groups (
    group_id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    owner_id INTEGER REFERENCES users(user_id)
);

-- Events table
CREATE TABLE IF NOT EXISTS events (
    event_id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    event_date DATE NOT NULL
);


-----------------------
-- Relational Tables --
-----------------------

-- Group Members table
CREATE TABLE IF NOT EXISTS group_members (
    group_id INTEGER REFERENCES groups(group_id) ON DELETE CASCADE,
    member INTEGER REFERENCES users(user_id),
    date_joined DATE NOT NULL,
    PRIMARY KEY (group_id, member)
);

-- Event Signups table
CREATE TABLE IF NOT EXISTS event_signups (
    event_id INTEGER REFERENCES events(event_id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(user_id),
    paid INTEGER NOT NULL CHECK (paid IN (0,1)), -- SQLite does not have boolean. Use 0 for False and 1 for True
    PRIMARY KEY (event_id, user_id)
);