import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: "libs/shared/graphql/*.graphql",
  documents: ["apps/frontend/src/**/*.tsx"],
  ignoreNoDocuments: true,
  generates: {
    "apps/frontend/src/graphql/": {
      preset: "client",
      presetConfig: {
        gqlTagName: "gql",
      },
      plugins: []
    },
  }
};

export default config;
