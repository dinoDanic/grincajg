defmodule GrincajgApiWeb.DefaultController do
  use GrincajgApiWeb, :controller

  def index(conn, _params) do
    text(conn, "grincajg_api is live - #{Mix.env()}")
  end
end
