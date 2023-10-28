ALTER TABLE organizations
DROP CONSTRAINT IF EXISTS fk_organizations_admin_user;

ALTER TABLE organizations
DROP CONSTRAINT IF EXISTS fk_organizations_user;

ALTER TABLE users
DROP CONSTRAINT IF EXISTS fk_users_organization;
