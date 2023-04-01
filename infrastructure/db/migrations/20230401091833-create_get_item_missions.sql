
-- +migrate Up
create table if not exists get_item_missions(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  item_id bigint not null references items(id),
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id),
  primary key(id)
);

comment on table get_item_missions is '特定のアイテム獲得ミッション達成条件';
comment on column get_item_missions.mission_id is 'ミッションID';
comment on column get_item_missions.item_id is '獲得必要アイテムID';
-- +migrate Down
drop table if exists get_item_missions;