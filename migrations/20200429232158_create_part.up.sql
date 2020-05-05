CREATE TABLE part_manufacturer(
        id varchar(100) NOT NULL PRIMARY KEY,
        name varchar(1000) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE part(
        id varchar(100) NOT NULL PRIMARY KEY,
        mnf_id varchar(100) REFERENCES part_manufacturer(id),
        vendor_code varchar(1000),
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        deleted_at TIMESTAMPTZ

);

--CREATE INDEX CONCURRENTLY mnfid_index ON part (mnf_id);

CREATE OR REPLACE RULE delete_part AS
  ON DELETE TO part
  -- WHERE old.deleted_at IS NULL
  DO INSTEAD
    UPDATE part SET deleted_at = NOW()
    WHERE part.id = old.id;
