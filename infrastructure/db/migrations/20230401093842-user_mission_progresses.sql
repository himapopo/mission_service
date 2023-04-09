
-- +migrate Up
create table if not exists user_mission_progresses(
  id bigserial not null,
  user_mission_id bigint not null references user_missions(id),
  progress_value bigint not null default 0,
  last_progress_updated_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  primary key(id)
);

comment on table user_mission_progresses is 'ユーザーのミッション毎の進捗';
comment on column user_mission_progresses.user_mission_id is 'ユーザーID';
comment on column user_mission_progresses.progress_value is '達成条件に関する現在の値';
comment on column user_mission_progresses.last_progress_updated_at is '最終進捗更新日時';
-- +migrate Down
drop table if exists user_mission_progresses;