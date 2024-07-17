CREATE TABLE IF NOT EXISTS story_tags (
    story_id UUID REFERENCES stories(id),
    tag VARCHAR(50),
    PRIMARY KEY (story_id, tag)
);