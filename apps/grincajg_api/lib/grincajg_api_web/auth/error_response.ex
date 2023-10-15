defmodule GrincajgApiWeb.Auth.ErrorResponse.Unauthorized do
  defexception message: "Unauthorized", plug_status: 401
end

defmodule GrincajgApiWeb.Auth.ErrorResponse.Forbidden do
  defexception message: "You do not have access", plug_status: 403
end
