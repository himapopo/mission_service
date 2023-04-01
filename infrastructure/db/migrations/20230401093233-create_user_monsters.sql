
-- +migrate Up
create table if not exists user_monsters(
  id bigserial not null,
  user_id bigint not null references users(id),
  monster_id bigint not null references monsters(id),
  level integer not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  primary key(id)
);

comment on table user_monsters is 'ユーザーの所有モンスター';
comment on column user_monsters.user_id is 'ユーザーID';
comment on column user_monsters.monster_id is 'モンスターID';
-- +migrate Down
drop table if exists user_monsters;