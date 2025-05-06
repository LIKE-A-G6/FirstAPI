ALTER TABLE tasks
    ADD COLUMN user_id INTEGER NOT NULL;

ALTER TABLE tasks
    ADD CONSTRAINT fk_user_task
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE;
