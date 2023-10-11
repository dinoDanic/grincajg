defmodule Grincajg.Repo do
  use Ecto.Repo,
    otp_app: :grincajg,
    adapter: Ecto.Adapters.Postgres
end
