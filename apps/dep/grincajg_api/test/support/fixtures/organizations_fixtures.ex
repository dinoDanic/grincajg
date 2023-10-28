defmodule GrincajgApi.OrganizationsFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `GrincajgApi.Organizations` context.
  """

  @doc """
  Generate a organization.
  """
  def organization_fixture(attrs \\ %{}) do
    {:ok, organization} =
      attrs
      |> Enum.into(%{

      })
      |> GrincajgApi.Organizations.create_organization()

    organization
  end
end
