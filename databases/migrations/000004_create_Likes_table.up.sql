CREATE TABLE IF NOT EXISTS likes (
    user_id UUID REFERENCES users(id),
    story_id UUID REFERENCES stories(id),
    liked_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, story_id)
);
