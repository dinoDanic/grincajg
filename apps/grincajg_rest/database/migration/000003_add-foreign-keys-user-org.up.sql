ALTER TABLE users
ADD CONSTRAINT fk_users_organization
FOREIGN KEY (organization_id)
REFERENCES organizations (id);

ALTER TABLE organizations
ADD CONSTRAINT fk_organizations_admin_user
FOREIGN KEY (admin_user_id)
REFERENCES users (id);
