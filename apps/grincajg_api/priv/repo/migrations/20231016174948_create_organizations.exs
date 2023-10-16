defmodule GrincajgApi.Repo.Migrations.CreateOrganizations do
  use Ecto.Migration

  def change do
    create table(:organizations) do
      add :name, :string, null: false
      add :address, :string
      add :account_id, references(:accounts, on_delete: :delete_all)

      timestamps(type: :utc_datetime)
    end

    create index(:organizations, [:account_id])
  end
end
