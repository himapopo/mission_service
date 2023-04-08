
-- +migrate Up
create table if not exists monster_level_up_missions(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  monster_id bigint not null references monsters(id),
  level integer not null,
  monster_count bigint not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id),
  primary key(id)
);

comment on table monster_level_up_missions is '特定のモンスターのレベルアップミッション達成条件';
comment on column monster_level_up_missions.mission_id is 'ミッションID';
comment on column monster_level_up_missions.monster_id is '対象モンスターID';
comment on column monster_level_up_missions.level is '対象のレベル';
comment on column monster_level_up_missions.monster_count is 'レベルアップ必要数';
-- +migrate Down
drop table if exists monster_level_up_missions;