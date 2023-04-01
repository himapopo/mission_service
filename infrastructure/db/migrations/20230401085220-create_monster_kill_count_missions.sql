
-- +migrate Up
create table if not exists monster_kill_count_missions(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  kill_count bigint not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id),
  primary key(id)
);

comment on table monster_kill_count_missions is 'モンスター討伐数ミッション達成条件';
comment on column monster_kill_count_missions.mission_id is 'ミッションID';
comment on column monster_kill_count_missions.kill_count is '討伐必要モンスター数';
-- +migrate Down
drop table if exists monster_kill_count_missions;