
-- +migrate Up
create table if not exists mission_releases(
  id bigserial not null,
  complete_mission_id bigint not null references missions(id),
  release_mission_id bigint not null references missions(id),
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  primary key(id)
);

comment on table mission_releases is 'ミッション解放管理';
comment on column mission_releases.complete_mission_id is '達成必要なミッションID';
comment on column mission_releases.release_mission_id is '解放されるミッションID';
-- +migrate Down
drop table if exists mission_releases;