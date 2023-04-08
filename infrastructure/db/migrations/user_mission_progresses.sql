
-- +migrate Up
create table if not exists user_mission_progresses(
  id bigserial not null,
  user_mission_id bigint not null references user_missions(id),
  mission_id bigint not null references missions(id),
  progress_value bigint not null default 0,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(user_mission_id, mission_id),
  primary key(id)
);

comment on table user_mission_progresses is 'ユーザーのミッション毎の進捗';
comment on column user_mission_progresses.user_mission_id is 'ユーザーID';
comment on column user_mission_progresses.mission_id is 'ミッションID';
comment on column user_mission_progresses.progress_value is '達成条件に関する現在の値';
-- +migrate Down
drop table if exists user_mission_progresses;