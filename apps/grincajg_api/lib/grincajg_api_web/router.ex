defmodule GrincajgApiWeb.Router do
  use GrincajgApiWeb, :router
  use Plug.ErrorHandler

  def handle_errors(conn, %{reason: %Phoenix.Router.NoRouteError{message: message}}) do
    conn |> json(%{errors: message}) |> halt()
  end

  def handle_errors(conn, %{reason: %{message: message}}) do
    conn |> json(%{errors: message}) |> halt()
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  pipeline :auth do
    plug GrincajgApiWeb.Auth.Pipeline
  end

  scope "/api", GrincajgApiWeb do
    pipe_through :api
    get "/", DefaultController, :index
    post "/accounts/create", AccountController, :create
    post "/accounts/sign_in", AccountController, :sign_in
  end

  scope "/api", GrincajgApiWeb do
    pipe_through [:api, :auth]
    get "/accounts/by_id/:id", AccountController, :show
  end
end
