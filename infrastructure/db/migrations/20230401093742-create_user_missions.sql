
-- +migrate Up
create table if not exists user_missions(
  id bigserial not null,
  user_id bigint not null references users(id),
  mission_id bigint not null references missions(id),
  completed_at timestamp with time zone null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(user_id, mission_id),
  primary key(id)
);

comment on table user_missions is 'ユーザーのミッション達成状況';
comment on column user_missions.user_id is 'ユーザーID';
comment on column user_missions.mission_id is 'ミッションID';
comment on column user_missions.completed_at is 'ミッション完了日';
-- +migrate Down
drop table if exists user_missions;