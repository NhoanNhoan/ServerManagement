create database ServerManagement;

use ServerManagement;

create table if not exists STATUS_ROW (
	id int,
	description varchar(20),

	primary key(id)
);

create table if not exists DC (
	id varchar(6),
	description nvarchar(20),

	primary key(id)
);

create table if not exists RACK (
	id varchar(6),
	description nvarchar(20),

	primary key(id)
);

create table if not exists RACK_UNIT(
	id varchar(6),
	description nvarchar(20),

	primary key(id)
);

create table if not exists IP_NET(
	id varchar(6),
	value varchar(12),
	id_STATUS_ROW int,

	primary key(id),
	foreign key(id_STATUS_ROW) references STATUS_ROW(id)
);

create table if not exists SERVER_STATUS (
	id varchar(6),
	description varchar(20),

	primary key(id)
);

create table if not exists ERROR_STATUS(
	id varchar(6),
	description nvarchar(20),

	primary key(id)
);

create table if not exists PERSON (
	id varchar(6),
	name nvarchar(50),

	primary key(id)
);

create table if not exists PARAMETER (
	
);

create table if not exists PORT_TYPE (
	id varchar(4),
	description nvarchar(10),

	primary key(id)
);

create table if not exists CABLE_TYPE (
	id varchar(6),
	name nvarchar(12),
	sign_port varchar(5),

	primary key(id)
);

create table if not exists SERVER (
	id varchar(12),
	id_DC varchar(6),
	id_RACK varchar(6),
	id_U_start varchar(6),
	id_U_end varchar(6),
	num_disks int,
	marker varchar(6),
	id_PORT_TYPE varchar(4),
	serial_number varchar(20),
	id_STATUS_ROW int,

	primary key(id),
	foreign key(id_DC) references DC(id),
	foreign key(id_RACK) references RACK(id),
	foreign key(id_U_start) references RACK_UNIT(id),
	foreign key(id_U_end) references RACK_UNIT(id),
	foreign key(id_PORT_TYPE) references PORT_TYPE(id),
	foreign key(id_STATUS_ROW) references STATUS_ROW(id)
);

create table if not exists SERVER_EVENT (
	id varchar(6) ,
	id_SERVER varchar(12),
	event nvarchar(500),
	occur_at date,
	id_STATUS_ROW int,

	primary key(id),
	foreign key(id_server) references SERVER(id),
	foreign key(id_STATUS_ROW) references STATUS_ROW(id)
);

create table if not exists SERVICES (
	id varchar(6),
	id_SERVER varchar(12),
	services varchar(50),

	primary key(id),
	foreign key(id_SERVER) references SERVER(id)
);

create table if not exists IP_SERVER (
	id_SERVER varchar(12),
	id_IP_NET varchar(6),
	ip_host int,
	id_STATUS_ROW int,

	foreign key(id_SERVER) references SERVER(id),
	foreign key(id_IP_NET) references IP_NET(id),
	primary key(id_SERVER, id_IP_NET, ip_host)
);

create table if not exists SWITCH (
	id varchar(10),
	name varchar(20),
	id_DC varchar(6),
	id_RACK varchar(6),
	id_U_start varchar(6),
	id_U_end varchar(6),
	maximum_port int,
	id_STATUS_ROW int,

	primary key(id),
	foreign key(id_DC) references DC(id),
	foreign key(id_RACK) references RACK(id),
	foreign key(id_U_start) references RACK_UNIT(id),
	foreign key(id_U_end) references RACK_UNIT(id),
	foreign key(id_STATUS_ROW) references STATUS_ROW(id)
);

create table if not exists IP_SWITCH (
	id_SWITCH varchar(10),
	id_IP_NET varchar(6),
	ip_host int,

	foreign key (id_SWITCH) references SWITCH(id),
	foreign key (id_IP_NET) references IP_NET(id),
	primary key(id_SWITCH, id_IP_NET, ip_host)
);

create table if not exists SWITCH_CONNECTION (
	id varchar(6) primary key,
	id_SERVER varchar(12),
	id_SWITCH varchar(10),
	id_CABLE_TYPE varchar(6),

	foreign key (id_SERVER) references SERVER(id),
	foreign key (id_SWITCH )references SWITCH(id),
	foreign key (id_CABLE_TYPE) references CABLE_TYPE(id)
);

create table if not exists SWITCH_CONNECTION_PORT (
	id_SWITCH varchar(10),
	sv_port int,
	switch_port int,

	foreign key (id_SWITCH) references SWITCH(id),
	primary key(id_SWITCH, sv_port, switch_port)
);

create table if not exists ERROR (
	id varchar(6),
	summary nvarchar(50),
	description nvarchar(500),
	solution nvarchar(1000),
	occurs date,
	id_SERVER varchar(12),
	id_ERROR_STATUS varchar(6),
	id_STATUS_ROW int,

	primary key(id),
	foreign key (id_SERVER) references SERVER(id),
	foreign key (id_ERROR_STATUS) references ERROR_STATUS(id),
	foreign key (id_STATUS_ROW) references STATUS_ROW(id)
);

create table if not exists USER (
	id varchar(10),
	username varchar(20),
	pass varchar(50),
	id_STATUS_ROW int,

	primary key(id),
	foreign key(id_STATUS_ROW) references STATUS_ROW(id)
);

create table if not exists CURATOR (
	id_PERSON varchar(6),
	id_ERROR varchar(6),
	id_STATUS_ROW int,

	foreign key (id_PERSON) references PERSON(id),
	foreign key (id_ERROR) references ERROR(id),
	foreign key (id_STATUS_ROW) references STATUS_ROW(id),
	primary key(id_PERSON, id_ERROR)
);
