
-- +migrate Up
create table if not exists user_monster_kill_histories(
  id bigserial not null,
  user_id bigint not null references users(id),
  monster_id bigint not null references monsters(id),
  killed_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  primary key(id)
);

comment on table user_monster_kill_histories is 'ユーザーのモンスター討伐履歴';
comment on column user_monster_kill_histories.user_id is 'ユーザーID';
comment on column user_monster_kill_histories.monster_id is 'モンスターID';
comment on column user_monster_kill_histories.killed_at is '討伐日';
-- +migrate Down
drop table if exists user_monster_kill_histories;