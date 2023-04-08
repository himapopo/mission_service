
-- +migrate Up
create table if not exists monster_kill_missions(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  monster_id bigint not null references monsters(id),
  monster_count bigint not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id),
  primary key(id)
);

comment on table monster_kill_missions is '特定のモンスター討伐ミッション達成条件';
comment on column monster_kill_missions.mission_id is 'ミッションID';
comment on column monster_kill_missions.monster_id is '対象モンスターID';
comment on column monster_kill_missions.monster_count is '討伐必要数';
-- +migrate Down
drop table if exists monster_kill_missions;