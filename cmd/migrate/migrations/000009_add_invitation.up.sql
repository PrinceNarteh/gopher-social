CREATE TABLE IF NOT EXISTS user_invitations (
  user_id bigint NOT NULL,
  token bytea PRIMARY KEY
);

