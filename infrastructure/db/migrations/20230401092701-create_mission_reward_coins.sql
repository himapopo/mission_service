
-- +migrate Up
create table if not exists mission_reward_coins(
  id bigserial not null,
  mission_id bigint not null references missions(id),
  coin_count bigint not null,
  updated_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  unique(mission_id, coin_count),
  primary key(id)
);

comment on table mission_reward_coins is 'ミッションの報酬コイン';
comment on column mission_reward_coins.mission_id is 'ミッションID';
comment on column mission_reward_coins.coin_count is '付与するコイン数';
-- +migrate Down
drop table if exists mission_reward_coins;