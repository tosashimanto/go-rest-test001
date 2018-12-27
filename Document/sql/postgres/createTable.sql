-- PostgreSQL

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;



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


-- アカウントテーブル
CREATE TABLE account (
    id bigserial PRIMARY KEY,
    role_id bigint NOT NULL,                      -- ロールID
    name varchar(40)                              -- 名前
    firebase_id varchar(28) UNIQUE,               -- FirebaseID
    created_by bigint NOT NULL REFERENCES account(id),
    created_at timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_by bigint NOT NULL REFERENCES account(id),
    updated_at timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc')
);
COMMENT ON TABLE account IS                'アカウントテーブル';
COMMENT ON COLUMN account.role_id IS       'ロールID';
COMMENT ON COLUMN account.operator_id IS   '作業員ID';
COMMENT ON COLUMN account.firebase_id IS   'FirebaseID';

-- ロールテーブル
CREATE TABLE role (
    id bigserial PRIMARY KEY,
    role_name varchar(100) NOT NULL,  -- ロール名
    created_by bigint NOT NULL REFERENCES account(id),
    created_at timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_by bigint NOT NULL REFERENCES account(id),
    updated_at timestamp with time zone NOT NULL DEFAULT (now() at time zone 'utc')
);
COMMENT ON TABLE role IS             'ロールテーブル';
COMMENT ON COLUMN role.role_name IS  'ロール名';

