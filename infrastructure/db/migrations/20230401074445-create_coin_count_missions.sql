
-- +migrate Up
create table if not exists coin_count_missions(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  coin_count bigint not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id),
  primary key(id)
);

comment on table coin_count_missions is '獲得コイン数ミッション達成条件';
comment on column coin_count_missions.mission_id is 'ミッションID';
comment on column coin_count_missions.coin_count is '獲得必要コイン数';
-- +migrate Down
drop table if exists coin_count_missions;