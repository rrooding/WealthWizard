/* eslint-disable */
import * as types from './graphql';
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  query CurrentlyHeldSecuritiesQuery {\n    ...SecurityItems_QueryFragment\n  }\n": types.CurrentlyHeldSecuritiesQueryDocument,
    "\n  fragment SecurityItems_QueryFragment on Query {\n    securities {\n      isin\n      ...SecurityItem_SecurityFragment\n    }\n  }\n": types.SecurityItems_QueryFragmentFragmentDoc,
    "\n  fragment SecurityItem_SecurityFragment on Security {\n    isin\n    amount\n    averagePrice {\n      ...MoneyItem_MoneyFragment\n    }\n  }\n": types.SecurityItem_SecurityFragmentFragmentDoc,
    "\n  fragment MoneyItem_MoneyFragment on Money {\n    amount\n    currency\n  }\n": types.MoneyItem_MoneyFragmentFragmentDoc,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query CurrentlyHeldSecuritiesQuery {\n    ...SecurityItems_QueryFragment\n  }\n"): (typeof documents)["\n  query CurrentlyHeldSecuritiesQuery {\n    ...SecurityItems_QueryFragment\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment SecurityItems_QueryFragment on Query {\n    securities {\n      isin\n      ...SecurityItem_SecurityFragment\n    }\n  }\n"): (typeof documents)["\n  fragment SecurityItems_QueryFragment on Query {\n    securities {\n      isin\n      ...SecurityItem_SecurityFragment\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment SecurityItem_SecurityFragment on Security {\n    isin\n    amount\n    averagePrice {\n      ...MoneyItem_MoneyFragment\n    }\n  }\n"): (typeof documents)["\n  fragment SecurityItem_SecurityFragment on Security {\n    isin\n    amount\n    averagePrice {\n      ...MoneyItem_MoneyFragment\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment MoneyItem_MoneyFragment on Money {\n    amount\n    currency\n  }\n"): (typeof documents)["\n  fragment MoneyItem_MoneyFragment on Money {\n    amount\n    currency\n  }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;
