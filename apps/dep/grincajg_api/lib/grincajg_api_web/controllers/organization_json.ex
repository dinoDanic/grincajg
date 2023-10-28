defmodule GrincajgApiWeb.OrganizationJSON do
  alias GrincajgApi.Organizations.Organization

  @doc """
  Renders a list of organizations.
  """
  def index(%{organizations: organizations}) do
    %{data: for(organization <- organizations, do: data(organization))}
  end

  @doc """
  Renders a single organization.
  """
  def show(%{organization: organization}) do
    data(organization)
  end

  def organization(%{organization: organization}) do
    data(organization)
  end

  defp data(%Organization{} = organization) do
    %{
      id: organization.id,
      name: organization.name,
      address: organization.address
    }
  end
end
