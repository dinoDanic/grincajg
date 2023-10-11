defmodule GrincajgApiWeb.Router do
  use GrincajgApiWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", GrincajgApiWeb do
    pipe_through :api
  end
end
