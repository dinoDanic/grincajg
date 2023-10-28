defmodule GrincajgApiWeb.UserController do
  use GrincajgApiWeb, :controller

  alias GrincajgApi.Users
  alias GrincajgApi.Users.User

  import GrincajgApiWeb.Auth.AuthorizedPlug

  # use PhoenixSwagger

  plug :is_authorized when action in [:update, :delete]
  action_fallback GrincajgApiWeb.FallbackController

  def index(conn, _params) do
    users = Users.list_users()
    render(conn, :index, users: users)
  end

  def create(conn, %{"user" => user_params}) do
    with {:ok, %User{} = user} <- Users.create_user(user_params) do
      conn
      |> put_status(:created)
      |> render(:show, user: user)
    end
  end

  def show(conn, %{"id" => id}) do
    user = Users.get_user!(id)
    render(conn, :show, user: user)
  end

  def update(conn, %{"user" => user_params}) do
    with {:ok, %User{} = user} <- Users.update_user(conn.assigns.account.user, user_params) do
      render(conn, :show, user: user)
    end
  end

  def delete(conn, %{"id" => id}) do
    user = Users.get_user!(id)

    with {:ok, %User{}} <- Users.delete_user(user) do
      send_resp(conn, :no_content, "")
    end
  end

  # def swagger_definitions do
  #   %{
  #     User:
  #       swagger_schema do
  #         title("User")
  #         description("A user of the app")
  #
  #         properties do
  #           id(:integer, "User ID")
  #           name(:string, "User name", required: true)
  #           email(:string, "Email address", format: :email, required: true)
  #           inserted_at(:string, "Creation timestamp", format: :datetime)
  #           updated_at(:string, "Update timestamp", format: :datetime)
  #         end
  #
  #         example(%{
  #           id: 123,
  #           name: "Joe",
  #           email: "joe@gmail.com"
  #         })
  #       end,
  #     UserRequest:
  #       swagger_schema do
  #         title("UserRequest")
  #         description("POST body for creating a user")
  #         property(:user, Schema.ref(:User), "The user details")
  #       end,
  #     UserResponse:
  #       swagger_schema do
  #         title("UserResponse")
  #         description("Response schema for single user")
  #         property(:data, Schema.ref(:User), "The user details")
  #       end,
  #     UsersResponse:
  #       swagger_schema do
  #         title("UsersReponse")
  #         description("Response schema for multiple users")
  #         property(:data, Schema.array(:User), "The users details")
  #       end
  #   }
  # end
end
