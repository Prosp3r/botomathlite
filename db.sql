CREATE TABLE botomath.users (
	id BIGSERIAL not null,
	email VARCHAR(50),
	datejoined TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
	lastlogin TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
	PRIMARY KEY (id, email));
	
create table botomath.robots (
	id BIGSERIAL not null,
	botname VARCHAR(50),
	robotype VARCHAR(50),
	datecreated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
	primary key (id));
	
create table botomath.robotype (
	ID BIGSERIAL not null,
	roboname VARCHAR(50),
	skill VARCHAR(200),
	primary key (id));
	
create table botomath.tasktype (
	id BIGSERIAL not null,
	description VARCHAR(50),
	eta int,
	primary key (id));
