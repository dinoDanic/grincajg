defmodule GrincajgApiWeb.AccountJSON do
  alias GrincajgApiWeb.OrganizationJSON
  alias Hex.API.Key.Organization
  alias GrincajgApiWeb.UserJSON
  alias GrincajgApi.Accounts.Account

  @doc """
  Renders a list of accounts.
  """
  def index(%{accounts: accounts}) do
    %{data: for(account <- accounts, do: data(account))}
  end

  @doc """
  Renders a single account.
  """
  def show(%{account: account}) do
    %{data: data(account)}
  end

  def render("account_token.json", %{account: account, token: token}) do
    %{
      id: account.id,
      email: account.email,
      token: token
    }
  end

  def full_account(%{account: account}) do
    %{
      id: account.id,
      email: account.email,
      user: UserJSON.user(%{user: account.user}),
      organization: OrganizationJSON.organization(%{organization: account.organization}) 
    }
  end

  defp data(%Account{} = account) do
    %{
      id: account.id,
      email: account.email,
      hash_password: account.hash_password
    }
  end
end
