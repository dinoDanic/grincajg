defmodule GrincajgApiWeb.OrganizationController do
  use GrincajgApiWeb, :controller
  use PhoenixSwagger

  alias GrincajgApi.Organizations
  alias GrincajgApi.Organizations.Organization
  import GrincajgApiWeb.Auth.AuthorizedPlug

  action_fallback GrincajgApiWeb.FallbackController

  def index(conn, _params) do
    organizations = Organizations.list_organizations()
    render(conn, :index, organizations: organizations)
  end

  def create(conn, %{"organization" => organization_params}) do
    account = conn.assigns.account

    params =
      %{
        "account_id" => account.id
      }
      |> Map.merge(organization_params)

    with {:ok, %Organization{} = organization} <-
           Organizations.create_organization(params) do
      conn
      |> put_status(:created)
      |> render(:show, organization: organization)
    end
  end

  def show(conn, %{"id" => id}) do
    organization = Organizations.get_organization!(id)
    render(conn, :show, organization: organization)
  end

  def update(conn, %{"id" => id, "organization" => organization_params}) do
    organization = Organizations.get_organization!(id)

    with {:ok, %Organization{} = organization} <-
           Organizations.update_organization(organization, organization_params) do
      render(conn, :show, organization: organization)
    end
  end

  def delete(conn, %{"id" => id}) do
    organization = Organizations.get_organization!(id)

    with {:ok, %Organization{}} <- Organizations.delete_organization(organization) do
      send_resp(conn, :no_content, "")
    end
  end

  # SWAGER

  # def swagger_definitions do
  #   %{
  #     Organization:
  #       swagger_schema do
  #         title("Organization")
  #         description("Organization")
  #
  #         properties do
  #           id(:string, "id")
  #           email(:string, "email")
  #         end
  #
  #         example(%{
  #           id: 1,
  #           token:
  #             "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJncmluY2FqZ19hcGkiLCJleHAiOjE2OTc2NDQ2MDEsImlhdCI6MTY5NzYzNzQwMSwiaXNzIjoiZ3JpbmNhamdfYXBpIiwianRpIjoiODAyOTFiNjItNTUwYi00ZTEyLTgwZmQtMjg3Y2MxOGZmNGI2IiwibmJmIjoxNjk3NjM3NDAwLCJzdWIiOiIyIiwidHlwIjoiYWNjZXNzIn0.XjLNL_ToTqxbsZAd6lJ7GPWSsL_KIUuDSOSBkQf83yzq7IqreaISrfE-_SIBdGc-q6SOLbIA366gU4lOD1hzaA",
  #           email: "email@email.com"
  #         })
  #       end
  #   }
  # end
  #
  # swagger_path :create do
  #   post("/accounts/sign_in")
  #   summary("List all recorded activities")
  #   description("List all recorded activities")
  #
  #   parameter(:email, :query, :string, "email", required: true, default: "vazin@kita.com")
  #   parameter(:hash_password, :query, :string, "passwordk", required: true, default: "1")
  #
  #   response(200, "Ok", Schema.ref(:Account))
  #   response(401, "Wrong credentials")
  # end
end
