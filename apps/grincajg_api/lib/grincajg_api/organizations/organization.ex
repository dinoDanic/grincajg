defmodule GrincajgApi.Organizations.Organization do
  use Ecto.Schema
  import Ecto.Changeset

  schema "organizations" do
    field :name, :string
    field :address, :string

    belongs_to :account, GrincajgApi.Accounts.Account

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(organization, attrs) do
    organization
    |> cast(attrs, [:name, :address, :account_id])
    |> validate_required([])
  end
end
