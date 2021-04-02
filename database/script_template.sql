create table CPU (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table RAM (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table DISK (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table RAID (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table NIC (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table PSU (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table MANAGEMENT (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table WARRANTLY (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table CHASSIS (
  id varchar(6),
  information varchar(50),
  primary key(id)
);

create table CLUSTER_SERVER (
  id varchar(6),
  name varchar(20),
  primary key(id)
);


create table HARDWARE_CONFIG (
  id varchar(6),
  chassis_id varchar(6),
  cluster_server_id varchar(6),
  primary key(id)
);
	hardware_id varchar(6),
	ram_id varchar(6),
	number_ram int,
	primary key(hardware_id, ram_id),
	foreign key (hardware_id) references HARDWARE_CONFIG(id),
	foreign key (ram_id) references RAM(id)
);

create table HARDWARE_DISK (
	hardware_id varchar(6),
	disk_id varchar(6),
	number_disk int,
	primary key(hardware_id, disk_id),
	foreign key (hardware_id) references HARDWARE_CONFIG(id),
	foreign key (disk_id) references DISK(id)
);

create table HARDWARE_RAID (
	hardware_id varchar(6),
	raid_id varchar(6),
	number_raid int,
	primary key(hardware_id, raid_id),
	foreign key (hardware_id) references HARDWARE_CONFIG(id),
	foreign key (raid_id) references RAID(id)
);

create table HARDWARE_NIC (
	hardware_id varchar(6),
	nic_id varchar(6),
	number_nic int,
	primary key(hardware_id, nic_id),
	foreign key (hardware_id) references HARDWARE_CONFIG(id),
	foreign key (nic_id) references NIC(id)
);

create table HARDWARE_PSU (
	hardware_id varchar(6),
	psu_id varchar(6),
	number_psu int,
	primary key(hardware_id, psu_id),
	foreign key (hardware_id) references HARDWARE_CONFIG(id),
	foreign key (psu_id) references PSU(id)
);

create table HARDWARE_MNT (
	hardware_id varchar(6),
	mnt_id varchar(6),
	number_mnt int,
	primary key(hardware_id, mnt_id),
	foreign key (hardware_id) references HARDWARE_CONFIG(id),
	foreign key (mnt_id) references MANAGEMENT(id)
);

create table HARDWARE_WARRANTLY (
	hardware_id varchar(6),
	warrantly_id varchar(6),
	number_warrantly int,
	primary key(hardware_id, warrantly_id),
	foreign key (hardware_id) references HARDWARE_CONFIG(id),
	foreign key (warrantly_id) references WARRANTLY(id)
);
