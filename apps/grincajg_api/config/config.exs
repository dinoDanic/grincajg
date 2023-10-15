# This file is responsible for configuring your application
# and its dependencies with the aid of the Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
import Config

config :grincajg_api,
  ecto_repos: [GrincajgApi.Repo],
  generators: [timestamp_type: :utc_datetime]

# Configures the endpoint
config :grincajg_api, GrincajgApiWeb.Endpoint,
  url: [host: "localhost"],
  adapter: Phoenix.Endpoint.Cowboy2Adapter,
  render_errors: [
    formats: [json: GrincajgApiWeb.ErrorJSON],
    layout: false
  ],
  pubsub_server: GrincajgApi.PubSub,
  live_view: [signing_salt: "cZP6gl8j"]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Guardian
config :grincajg_api, GrincajgApiWeb.Auth.Guardian,
  issuer: "grincajg_api",
  secret_key: "tCWqH9kRzzilziqeszd21uxy9mwoQYzzYneEr0k0JkqYXrj6mH1kwPca+xdvOqSW"

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

config :guardian, Guardian.DB,
  repo: GrincajgApi.Repo,
  schema_name: "guardian_tokens",
  sweep_interval: 60

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{config_env()}.exs"
