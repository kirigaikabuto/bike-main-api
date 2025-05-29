-- migrate:up
alter table users add column email text;

-- migrate:down
alter table users drop column email;