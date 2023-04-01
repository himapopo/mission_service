
-- +migrate Up
create table if not exists user_items(
  id bigserial not null,
  user_id bigint not null references users(id),
  item_id bigint not null references items(id),
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  primary key(id)
);

comment on table user_items is 'ユーザーの所有アイテム';
comment on column user_items.user_id is 'ユーザーID';
comment on column user_items.item_id is 'アイテムID';
-- +migrate Down
drop table if exists user_items;