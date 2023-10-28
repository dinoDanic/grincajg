# Script for populating the database. You can run it as:
#
#     mix run priv/repo/seeds.exs
#
# Inside the script, you can read and write to any of your
# repositories directly:
#
#     GrincajgApi.Repo.insert!(%GrincajgApi.SomeSchema{})
#
# We recommend using the bang functions (`insert!`, `update!`
# and so on) as they will fail if something goes wrong.


GrincajgApi.Repo.insert!(%GrincajgApi.Accounts.Account{
  email: "vazin@gmail.com",
  hash_password: "$2b$12$0stZDac46ewTe5uE5FY3GOrjkQTcNVlMfwIa0L28ykbJbbeFtGyIO",
  user: %GrincajgApi.Users.User{
    first_name: "kita",
    last_name: "me"
  }
})

GrincajgApi.Repo.insert!(%GrincajgApi.Organizations.Organization{
  name: "Org 1",
  account_id: 1
})
