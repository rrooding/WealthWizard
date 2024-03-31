import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: "libs/shared/graphql/*.graphql",
  documents: [
    "apps/frontend/src/**/*.tsx",
    "libs/web/**/*.tsx"
  ],
  ignoreNoDocuments: true,
  generates: {
    "libs/web/graphql/src/": {
      preset: "client",
      config: {
        useTypeImports: true,
        strictScalars: true,
        scalars: {
          Decimal: "number",
          DateTime: "Date",
        },
      },
      presetConfig: {
        fragmentMasking: { unmaskFunctionName: 'getFragmentData' },
        gqlTagName: "gql",
      },
      plugins: []
    },
  }
};

export default config;
