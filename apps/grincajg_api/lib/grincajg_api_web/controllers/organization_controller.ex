defmodule GrincajgApiWeb.OrganizationController do
  use GrincajgApiWeb, :controller
  use PhoenixSwagger

  alias Hex.API.Key.Organization
  alias GrincajgApi.Organizations
  alias GrincajgApi.Organizations.Organization

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

  def show_active(conn, %{}) do
    current_account = conn.assigns.account
    preload = Organizations.preload_organization(current_account)

    conn
    |> render(:show, organization: preload.organization)
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

  def swagger_definitions do
    %{
      Organization:
        swagger_schema do
          title("organization")
          description("organization")

          properties do
            id(:string)
            name(:string)
            address(:string)
          end

          example(%{
            id: 1,
            name: "organization example",
            address: "franc 1, 100 000 zamra"
          })
        end,
      OrganizationInput:
        swagger_schema do
          title("organization input")

          description(
            "arguments for creating a organization, a user can have only one organization"
          )

          properties do
            name(:string)
            address(:string)
          end

          example(%{
            organization: %{
              name: "organization example",
              address: "franc 1, 100 000 zamra"
            }
          })
        end
    }
  end

  swagger_path :create do
    post("/organization/create")
    summary("create a organization")
    description("")

    parameters do
      organization(:body, Schema.ref(:OrganizationInput), "Organization attributes",
        required: true
      )
    end

    security([%{Bearer: []}])

    response(200, "ok", Schema.ref(:Organization))
    response(401, "wrong credentials")
  end

  swagger_path :show_active do
    get("/organization/active")
    summary("Get active account organization")

    security([%{Bearer: []}])

    response(200, "Ok", Schema.ref(:Organization))
  end
end
