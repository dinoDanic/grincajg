defmodule GrincajgApiWeb.AccountController do
  use GrincajgApiWeb, :controller

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
      authorize_account(conn, account.email, account_params[~c"hash_password"])
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

  def show(conn, %{"id" => id}) do
    account = Accounts.get_me_account(id)
    render(conn, :full_account, account: account)
  end

  def me_account(conn, %{}) do
    account = conn.assigns.account

    conn
    |> put_status(:ok)

    render(conn, :render_me_account, account: account)
  end

  def me_organization(conn, %{}) do
    current_account = conn.assigns.account
    account_organization = Accounts.preload_organization(current_account)

    conn
    |> put_status(:ok)
    |> render(:render_me_organization, account: account_organization)
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
end
