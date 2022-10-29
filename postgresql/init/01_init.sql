-- 初期設定
create database todolist;

-- スキーマ作成
\c todolist
create schema todolistschema;

-- ロールの作成
create role todolist_user with login password 'passw0rd';
alter role todolist_user with createdb createrole;

-- 権限を追加
grant all privileges on schema todolistschema to todolist_user;

-- 作成したDBへ切り替える
\c todolist;
\c - todolist_user;
