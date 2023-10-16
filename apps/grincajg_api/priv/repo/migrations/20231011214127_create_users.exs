defmodule GrincajgApi.Repo.Migrations.CreateUsers do
  use Ecto.Migration

  def change do
    create table(:users) do
      add :first_name, :string
      add :last_name, :string
      add :gender, :string
      add :account_id, references(:accounts, on_delete: :delete_all)

      timestamps(type: :utc_datetime)
    end

    create index(:users, [:account_id])
  end
end
