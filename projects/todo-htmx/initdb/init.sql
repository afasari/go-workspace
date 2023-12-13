CREATE TABLE todos (
id serial not null,
todo varchar(255) not null,
done boolean not null,
PRIMARY KEY(id)
);