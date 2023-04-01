
-- +migrate Up
create table if not exists mission_reward_items(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  item_id bigint not null references items(id),
  item_count bigint not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id, item_id, item_count),
  primary key(id)
);

comment on table mission_reward_items is 'ミッションの報酬アイテム';
comment on column mission_reward_items.mission_id is 'ミッションID';
comment on column mission_reward_items.item_id is '付与するアイテムID';
comment on column mission_reward_items.item_count is '付与するアイテムの個数';
-- +migrate Down
drop table if exists mission_reward_items;