defmodule GrincajgApiWeb.Swagger.RandomController do
  use GrincajgApiWeb, :controller

  use PhoenixSwagger

  def swagger_definitions do
    %{
      User:
        swagger_schema do
          title("User")
          description("A user of the app")

          properties do
            id(:integer, "User ID")
            name(:string, "User name", required: true)
            email(:string, "Email address", format: :email, required: true)
            inserted_at(:string, "Creation timestamp", format: :datetime)
            updated_at(:string, "Update timestamp", format: :datetime)
          end

          example(%{
            id: 123,
            name: "Joe",
            email: "joe@gmail.com"
          })
        end,
      UserRequest:
        swagger_schema do
          title("UserRequest")
          description("POST body for creating a user")
          property(:user, Schema.ref(:User), "The user details")
        end,
      UserResponse:
        swagger_schema do
          title("UserResponse")
          description("Response schema for single user")
          property(:data, Schema.ref(:User), "The user details")
        end,
      UsersResponse:
        swagger_schema do
          title("UsersReponse")
          description("Response schema for multiple users")
          property(:data, Schema.array(:User), "The users details")
        end
    }
  end

  swagger_path :show do
    get("/users/{id}")
    summary("Retrieve a user")
    description("Get details of a specific user by ID")
    parameter(:id, :path, :integer, "ID of the user")
    response(200, "OK", Schema.ref(:UserResponse))
  end

  def show(conn, %{"id" => id}) do
    # Your implementation to retrieve and show a user
  end
end
