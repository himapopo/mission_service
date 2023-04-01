
-- +migrate Up
create table if not exists login_missions(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  login_count bigint not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id),
  primary key(id)
);

comment on table login_missions is 'ログイン回数ミッション達成条件';
comment on column login_missions.mission_id is 'ミッションID';
comment on column login_missions.login_count is 'ログイン回数';
-- +migrate Down
drop table if exists login_missions;