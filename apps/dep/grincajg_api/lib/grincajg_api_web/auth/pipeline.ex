defmodule GrincajgApiWeb.Auth.Pipeline do
  use Guardian.Plug.Pipeline,
    otp_app: :grincajg_api,
    module: GrincajgApiWeb.Auth.Guardian,
    error_handler: GrincajgApiWeb.Auth.GuardianErrorHandler

  plug Guardian.Plug.VerifySession
  plug Guardian.Plug.VerifyHeader
  plug Guardian.Plug.EnsureAuthenticated
  plug Guardian.Plug.LoadResource
end
