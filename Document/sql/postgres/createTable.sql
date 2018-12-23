-- PostgreSQL

DROP TABLE IF exists test002;

CREATE TABLE test002 (
    testID bigserial primary key,
    value integer null,
    name varchar(300) null,
    createdBy bigint not null,
    createdAt timestamp with time zone not null default (now() at time zone 'utc'),
    updatedBy bigint not null,
    updatedAt timestamp with time zone not null default (now() at time zone 'utc')
);
