\c todolist
\c - todolist_user

-- テーブルを作成
create table todolistschema.tasks (
	id serial unique primary key,
	name VARCHAR(100),
	discription VARCHAR(500),
	status VARCHAR(100)
);

-- テーブルにデータを挿入
insert into todolistschema.tasks values (
	1,
	'Sample task',
	'Test task',
	'In progress'
);
