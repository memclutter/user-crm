create table countries
(
	code varchar(2) not null
		constraint countries_pk
			primary key,
	name varchar(255) not null
);

create table users
(
	id serial not null,
	username varchar(255) not null,
	email varchar(255) not null,
	gender varchar(32) default NULL::character varying,
	birthday date,
	country_code varchar(2) default NULL::character varying
		constraint users_countries_code_fk
			references countries
				on update cascade on delete restrict
);

create unique index users_email_uindex
	on users (email);

create unique index users_id_uindex
	on users (id);

create unique index users_username_uindex
	on users (username);

create unique index countries_code_uindex
	on countries (code);

