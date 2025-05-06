ALTER TABLE tasks
DROP CONSTRAINT fk_user_task;

ALTER TABLE tasks
DROP COLUMN user_id;
