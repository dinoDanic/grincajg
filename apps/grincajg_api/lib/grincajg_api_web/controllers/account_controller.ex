defmodule GrincajgApiWeb.AccountController do
  use GrincajgApiWeb, :controller
  use PhoenixSwagger

  alias GrincajgApiWeb.Auth.ErrorResponse
  alias GrincajgApiWeb.Auth.Guardian
  alias GrincajgApi.{Accounts, Accounts.Account, Users, Users.User}

  import GrincajgApiWeb.Auth.AuthorizedPlug

  # WE DO THIS PLUG WHEN UPDATE OR DELETE FN ARE CALLED
  plug :is_authorized when action in [:update, :delete]

  action_fallback GrincajgApiWeb.FallbackController

  def index(conn, _params) do
    accounts = Accounts.list_accounts()
    render(conn, :index, accounts: accounts)
  end

  def create(conn, %{"account" => account_params}) do
    with {:ok, %Account{} = account} <- Accounts.create_account(account_params),
         {:ok, %User{} = _user} <- Users.create_user(account, account_params) do
      authorize_account(conn, account.email, account_params["hash_password"])
    end
  end

  def sign_in(conn, %{"email" => email, "hash_password" => hash_password}) do
    authorize_account(conn, email, hash_password)
  end

  def refresh_session(conn, %{}) do
    token = Guardian.Plug.current_token(conn)
    {:ok, account, new_token} = Guardian.authenticate(token)

    conn
    |> Plug.Conn.put_session(:account_id, account.id)
    |> put_status(:ok)
    |> render("account_token.json", %{account: account, token: new_token})
  end

  def sign_out(conn, %{}) do
    account = conn.assigns[:account]
    token = Guardian.Plug.current_token(conn)
    Guardian.revoke(token)

    conn
    |> Plug.Conn.clear_session()
    |> put_status(:ok)
    |> render("account_token.json", %{account: account, token: nil})
  end

  defp authorize_account(conn, email, hash_password) do
    case Guardian.authenticate(email, hash_password) do
      {:ok, account, token} ->
        conn
        |> Plug.Conn.put_session(:account_id, account.id)
        |> put_status(:ok)
        |> render("account_token.json", %{account: account, token: token})

      {:error, :unauthorized} ->
        raise ErrorResponse.Unauthorized, message: "Email or Password incorrect."
    end
  end

  def account_user(conn, %{}) do
    account = conn.assigns.account
    render(conn, :render_account_user, account: account)
  end

  def account_user_with_organization(conn, %{}) do
    current_account = conn.assigns.account
    account_organization = Accounts.preload_organization(current_account)

    conn
    |> put_status(:ok)
    |> render(:render_account_user_organization, account: account_organization)
  end

  def update(conn, %{"current_hash" => current_hash, "account" => account_params}) do
    account = conn.assigns.account

    case Guardian.validate_password(current_hash, account.hash_password) do
      true ->
        {:ok, account} = Accounts.update_account(account, account_params)
        render(conn, :show, account: account)

      false ->
        raise ErrorResponse.Unauthorized, message: "Passwod incorect"
    end
  end

  def delete(conn, %{"id" => id}) do
    account = Accounts.get_account!(id)

    with {:ok, %Account{}} <- Accounts.delete_account(account) do
      send_resp(conn, :no_content, "")
    end
  end

  # SWAGGER

  def swagger_definitions do
    %{
      Account:
        swagger_schema do
          title("Account")
          description("Account")

          properties do
            id(:string)
            email(:string)
          end

          example(%{
            id: 1,
            token: "Bearer token",
            email: "example@email.com"
          })
        end,
      AccountInput:
        swagger_schema do
          title("AccountInput")
          description("Arguments for creating account")

          properties do
            id(:string)
            email(:string)
          end

          example(%{account: %{email: "vazin@gmail.com", hash_password: "1"}})
        end
    }
  end

  swagger_path :create do
    post("/accounts/create")
    summary("create a account")
    # description("List all recorded activities")

    parameters do
      account(:body, Schema.ref(:AccountInput), "Account attributes", required: true)
    end

    response(200, "Ok", Schema.ref(:Account))
  end

  swagger_path :sign_in do
    post("/accounts/sign_in")
    summary("sign into account")
    # description("List all recorded activities")

    parameter(:email, :query, :string, "email", required: true, default: "vazin@gmail.com")
    parameter(:hash_password, :query, :string, "password", required: true, default: "1")

    response(200, "Ok", Schema.ref(:Account))
    response(401, "Wrong credentials")
  end

  swagger_path :sign_out do
    get("/accounts/sign_out")
    summary("sing out from account")

    security([%{Bearer: []}])

    response(200, "Ok", Schema.ref(:Account))
    response(401, "aunthenticated")
    response(402, "invalid_token")
  end

  swagger_path :refresh_session do
    get("/accounts/refresh_session")
    summary("refresh account session")

    security([%{Bearer: []}])

    parameters do
      account(:body, Schema.ref(:AccountInput), "Account attributes", required: true)
    end

    response(200, "Ok", Schema.ref(:Account))
  end

  swagger_path :account_user do
    get("/accounts/account_user")
    summary("get account info")

    security([%{Bearer: []}])

    response(200, "Ok", Schema.ref(:Account))
  end

  swagger_path :account_user_with_organization do
    get("/accounts/account_user_with_organization")
    summary("get organization from user")

    security([%{Bearer: []}])

    response(200, "Ok", Schema.ref(:Organization))
  end
end
