CREATE TABLE IF NOT EXISTS tenant (
    tenant_name_id VARCHAR(20) PRIMARY KEY,
    tenant_uuid UUID NOT NULL,
    created_at TIMESTAMPTZ
);

CREATE INDEX idx_tenant_uuid ON tenant (tenant_name_id);


CREATE TABLE IF NOT EXISTS tenant_user (
    tenant_name_id VARCHAR(20)
        REFERENCES tenant(tenant_name_id) ON DELETE CASCADE,
    tenant_user_id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    email TEXT
);

CREATE INDEX ON tenant_user (tenant_name_id);
ALTER TABLE tenant_user ADD CONSTRAINT id_name_unique UNIQUE(tenant_name_id,name);

ALTER TABLE tenant_user ADD CONSTRAINT id_mail_unique UNIQUE(tenant_name_id,email);

CREATE TABLE IF NOT EXISTS branches (
    tenant_name_id VARCHAR(20)
        REFERENCES tenant(tenant_name_id) ON DELETE CASCADE,
    branch_id UUID NOT NULL PRIMARY KEY,
    branch_name VARCHAR(20) NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX ON branches (tenant_name_id);


CREATE TABLE IF NOT EXISTS user_title (
    tenant_name_id VARCHAR(20)
        REFERENCES tenant(tenant_name_id) ON DELETE CASCADE,
    title_id UUID NOT NULL PRIMARY KEY,
    title_name VARCHAR(20) NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX ON user_title (tenant_name_id)
