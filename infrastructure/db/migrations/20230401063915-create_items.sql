
-- +migrate Up
create table if not exists items(
  id bigserial not null,
  name text not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(name),
  primary key(id)
);

comment on table items is 'アイテム';
comment on column items.name is '名前';
-- +migrate Down
drop table if exists items;