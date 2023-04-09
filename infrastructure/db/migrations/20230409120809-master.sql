
-- +migrate Up
-- モンスター
insert into monsters (name) values('モンスターA');
insert into monsters (name) values('モンスターB');

-- アイテム
insert into items (name) values('アイテムA');

-- ミッション
insert into missions (name, mission_type) values('特定のモンスターを倒す', 'normal');
insert into monster_kill_missions (mission_id, monster_id, monster_count) values(1, 1, 1);
insert into mission_reward_coins (mission_id, coin_count) values(1, 100);

insert into missions (name, mission_type) values('2000コイン貯まる', 'normal');
insert into coin_count_missions (mission_id, coin_count) values(2, 2000);
insert into mission_reward_items (mission_id, item_id, item_count) values(2, 1, 1);

insert into missions (name, mission_type) values('モンスターAのレベルが5になる', 'normal');
insert into monster_level_up_missions (mission_id, monster_id, level, monster_count) values(3, 1, 5, 1);
insert into mission_reward_coins (mission_id, coin_count) values(3, 100);

insert into missions (name, mission_type) values('レベル５以上のモンスターが２体', 'normal');
insert into monster_level_up_count_missions (mission_id, monster_count, level) values(4, 2, 5);
insert into mission_reward_coins (mission_id, coin_count) values(4, 100);


insert into missions (name, mission_type) values('任意のモンスターを10回倒す', 'weekly');
insert into monster_kill_count_missions (mission_id, kill_count) values(5, 10);
insert into mission_reward_coins (mission_id, coin_count) values(5, 100);

insert into missions (name, mission_type) values('ログイン', 'daily');
insert into login_missions (mission_id, login_count) values(6, 1);
insert into mission_reward_coins (mission_id, coin_count) values(6, 100);


insert into missions (name, mission_type) values('アイテムAを所有する', 'normal');
insert into get_item_missions (mission_id, item_id, item_count) values(7, 1, 1);
insert into mission_reward_coins (mission_id, coin_count) values(7, 100);

insert into mission_releases (complete_mission_id, release_mission_id) values (4, 7);
