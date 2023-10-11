defmodule GrincajgWeb.Router do
  use GrincajgWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", GrincajgWeb do
    pipe_through :api
  end
end
