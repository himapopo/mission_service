
-- +migrate Up
create table if not exists monsters_level_up_missions(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  monster_count bigint not null,
  level integer not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id),
  primary key(id)
);

comment on table monsters_level_up_missions is '一定レベル以上のモンスター獲得ミッション達成条件';
comment on column monsters_level_up_missions.mission_id is 'ミッションID';
comment on column monsters_level_up_missions.monster_count is '獲得必要モンスター数';
comment on column monsters_level_up_missions.level is '対象のレベル';
-- +migrate Down
drop table if exists monsters_level_up_missions;