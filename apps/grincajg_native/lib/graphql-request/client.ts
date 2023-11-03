import { GraphQLClient } from "graphql-request"

const endpoint = "https://grincajg-be-dev.dinosur.app/api"

export const _client = new GraphQLClient(endpoint)
