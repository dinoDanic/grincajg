defmodule GrincajgApi.Users.User do
  use Ecto.Schema
  import Ecto.Changeset

  schema "users" do
    field :first_name, :string
    field :last_name, :string
    field :gender, :string

    belongs_to :account, GrincajgApi.Accounts.Account

    timestamps(type: :utc_datetime)
  end

  @doc false
  def changeset(user, attrs) do
    user
    |> cast(attrs, [:account_id, :first_name, :last_name, :gender])
    |> validate_required([:account_id])
  end
end
