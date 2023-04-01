
-- +migrate Up
create table if not exists missions(
  id bigserial not null,
  name text not null,
  is_deleted boolean not null default false,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(name),
  primary key(id)
);

comment on table missions is 'ミッション';
comment on column missions.name is '名前';
comment on column missions.is_deleted is '削除フラグ';
-- +migrate Down
drop table if exists missions;