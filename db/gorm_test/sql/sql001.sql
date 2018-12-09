

DROP TABLE IF exists test001;
CREATE TABLE test001 (
    constructionID bigserial primary key,
    value integer null,
    name varchar(300) null
);

INSERT INTO test001 (
     value,
     name
) VALUES
(
    123456,
    'this is test001.'
);



DROP TABLE IF exists test002;
CREATE TABLE test002 (
    constructionID bigserial primary key,
    value integer null,
    name varchar(300) null,
    createdBy bigint not null,
    createdAt timestamp with time zone not null default (now() at time zone 'utc'),
    updatedBy bigint not null,
    updatedAt timestamp with time zone not null default (now() at time zone 'utc')
);

INSERT INTO test002 (
     value,
     name,
     createdBy,
     updatedBy
) VALUES
(
    123456,
    'this is test001.',
    99999999,
    99999999
);