import type { CodegenConfig } from "@graphql-codegen/cli"

const config: CodegenConfig = {
  overwrite: true,
  schema: "https://grincajg-be-dev.dinosur.app/api",
  // documents: "gql/**/*.graphql",
  documents: "*/**/*.tsx",
  generates: {
    "generated/": {
      preset: "client",
      plugins: [],
    },
  },
}

export default config
