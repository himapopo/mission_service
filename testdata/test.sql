-- ユーザー初期状態テストデータ
insert into users (name, coin_count, last_login_at) values('ユーザーA', 1800, '2023-04-07 01:00:00');

insert into user_monsters (user_id, monster_id, level) values(1, 1, 4);
insert into user_monsters (user_id, monster_id, level) values(1, 2, 7);


insert into user_missions (user_id, mission_id) values(1, 1);
insert into user_mission_progresses (user_mission_id) values(1);

insert into user_missions (user_id, mission_id) values(1, 2);
insert into user_mission_progresses (user_mission_id, progress_value) values(2, 1800);

insert into user_missions (user_id, mission_id) values(1, 3);
insert into user_mission_progresses (user_mission_id) values(3);


insert into user_missions (user_id, mission_id) values(1, 4);
insert into user_mission_progresses (user_mission_id, progress_value) values(4, 1);

insert into user_missions (user_id, mission_id) values(1, 5);
insert into user_mission_progresses (user_mission_id, progress_value, last_progress_updated_at) values(5, 8, '2023-04-05 02:00:00');

insert into user_missions (user_id, mission_id, completed_at) values(1, 6, '2023-04-07 01:00:00');
insert into user_mission_progresses (user_mission_id, progress_value, last_progress_updated_at) values(6, 1, '2023-04-07 01:00:00');