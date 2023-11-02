import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: "https://grincajg-be-dev.dinosur.app/query",
  documents: "gql/**/*.graphql",
  generates: {
    "generated/": {
      preset: "client",
      plugins: [],
    },
  },
};

export default config;
