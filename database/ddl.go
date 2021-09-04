// данный файл содержит команды для создания sql-таблиц

package database


//	ClientsDataDDL таблица с данными клиентов банка(физ.лиц)
const ClientsDataDDL = `create table if not exists client_Data (
	id serial primary key,
	name varchar(30) not null,
	surname text not null,
	login text not null unique,
	password text not null,
	phone text not null,
	locked boolean not null,
	creationDate date default current_timestamp,
	upDate date,
	deleted boolean
);`


//	MerchantsDataDDL - таблица с данными клиентов банка (юр.лиц)
const MerchantsDataDDL = `create table if not exists merchant_Data (
	id serial primary key,
	name varchar(30) not null,
	surname text not null,
	login text not null unique,
	password text not null,
	phone text not null,
	locked boolean not null,
	creationDate date default current_timestamp,
	upDate date,
	deleted boolean
);`

//	ManagersDataDDL - таблица с данными менеджеров банка
const ManagersDataDDL = `create table if not exists managers (
	id serial primary key,
	name varchar(30) not null,
	surname text not null,
	login text not null unique,
	password text not null,
	phone text not null,
	locked boolean not null,
	creationDate date default current_timestamp,
	upDate date,
	deleted boolean
);`

// ClientAccountsDDL - таблица аккаунтов клиентов банка (физ.лиц)
const ClientAccountsDDL = `create table if not exists Client_Accounts (
	id serial primary key,
	accNumber text not null,
	clientID integer,
	balance int, 
	swift text not null,
	LastVisitDate date not null,
	locked boolean,
	creationDate date default current_timestamp,
	upDate date,	
	deleted boolean
);`

// MerchantAccountsDDL - таблица аккаунтов клиентов банка (юр.лиц), занимающихся бизнесом.
const MerchantAccountsDDL = `create table if not exists Merchant_Accounts (
	id serial primary key,
	accNumber text not null,
	merchantID integer,
	balance int, 
	swift text not null,
	lastVisitDate date not null,
	locked boolean,
	creationDate date default current_timestamp,
	upDate date,	
	deleted boolean
);`

//	ServicesDDL - таблица онлайн услуг
const ServicesDDL = `create table if not exists services (
	ID serial primary key,
	serviceName text unique not null,
	merchantID integer,
	merchantName text not null,
	price integer not null,
	creationDate date default current_timestamp,
	upDate date,
	deleted boolean
);`

//	TransactionsDDL - таблица данных об осуществляемых трансакциях
const TransactionsDDL = `create table if not exists transactions (
	id serial primary key,
	payerAcc text not null,
	beneficiaryAcc text not null,
	transferAmount integer,
	creationDate date default current_timestamp,
	completed boolean
);`

//	atmsDDL - таблица с данными о банкоматах
const ATMsDDL = `create table if not exists ATMs(
	id serial primary key,
	address text not null,
	balance integer,
	status text,
	maxCashLimit integer,
	comissionFee boolean
);`

