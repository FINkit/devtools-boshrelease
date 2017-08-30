
drop table account_group_by_id_aud;
drop table account_group_members_audit;
drop table changes;

CREATE TABLE account_group_by_id_aud (
    added_by INT DEFAULT 0 NOT NULL,
    removed_by INT,
    removed_on TIMESTAMP NULL DEFAULT NULL,
    group_id INT DEFAULT 0 NOT NULL,
    include_uuid VARCHAR(255) BINARY DEFAULT '' NOT NULL,
    added_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(group_id,include_uuid,added_on)
);

CREATE TABLE account_group_members_audit (
    added_by INT DEFAULT 0 NOT NULL,
    removed_by INT,
    removed_on TIMESTAMP NULL DEFAULT NULL,
    account_id INT DEFAULT 0 NOT NULL,
    group_id INT DEFAULT 0 NOT NULL,
    added_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(account_id,group_id,added_on)
);

CREATE TABLE changes (
    change_key VARCHAR(60) BINARY DEFAULT '' NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_updated_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    owner_account_id INT DEFAULT 0 NOT NULL,
    dest_project_name VARCHAR(255) BINARY DEFAULT '' NOT NULL,
    dest_branch_name VARCHAR(255) BINARY DEFAULT '' NOT NULL,
    status CHAR(1) DEFAULT ' ' NOT NULL,
    current_patch_set_id INT DEFAULT 0 NOT NULL,
    subject VARCHAR(255) BINARY DEFAULT '' NOT NULL,
    topic VARCHAR(255) BINARY,
    original_subject VARCHAR(255) BINARY,
    submission_id VARCHAR(255) BINARY,
    assignee INT,
    note_db_state TEXT,
    row_version INT DEFAULT 0 NOT NULL,
    change_id INT DEFAULT 0 NOT NULL,
    PRIMARY KEY(change_id)
);

