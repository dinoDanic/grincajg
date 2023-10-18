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
    plug :fetch_session
  end

  pipeline :auth do
    plug GrincajgApiWeb.Auth.Pipeline
    plug GrincajgApiWeb.Auth.SetAccount
  end

  # PUBLIC ROUTE
  scope "/api", GrincajgApiWeb do
    pipe_through :api
    get "/", DefaultController, :index
    post "/accounts/create", AccountController, :create
    post "/accounts/sign_in", AccountController, :sign_in
  end

  # PROTECTED ROUTE
  scope "/api", GrincajgApiWeb do
    pipe_through [:api, :auth]
    get "/accounts/account_user", AccountController, :account_user
    get "/accounts/account_user_with_organization", AccountController, :account_user_with_organization
    get "/accounts/sign_out", AccountController, :sign_out
    get "/accounts/refresh_session", AccountController, :refresh_session
    post "/accounts/update", AccountController, :update
    put "/users/update", UserController, :update
    post "/organization/create", OrganizationController, :create
  end

  # SWAGER
  scope "/api/swagger" do
    forward "/", PhoenixSwagger.Plug.SwaggerUI,
      otp_app: :grincajg_api,
      swagger_file: "swagger.json"
  end

  def swagger_info do
    %{
      basePath: "/api",
      info: %{
        version: "1.0",
        title: "Grincajg"
      },
      securityDefinitions: %{
        Bearer: %{
          type: "apiKey",
          name: "Authorization",
          in: "header"
        }
      }
    }
  end
end
