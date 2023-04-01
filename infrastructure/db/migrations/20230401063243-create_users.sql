
-- +migrate Up
create table if not exists users(
  id bigserial not null,
  name text not null,
  email text not null,
  password text not null,
  coin_count bigint not null default 0,
  last_login_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(email, name),
  primary key(id)
);

comment on table users is 'ユーザー';
comment on column users.name is '名前';
comment on column users.email is 'メールアドレス';
comment on column users.password is '暗号化済みパスワード';
comment on column users.coin_count is '獲得コイン数';
-- +migrate Down
drop table if exists users;