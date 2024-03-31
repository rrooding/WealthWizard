/* eslint-disable */
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  DateTime: { input: Date; output: Date; }
  /**
   * A signed decimal number, which supports arbitrary precision and is serialized as a string.
   *
   * Example values: 29.99, 29.999.
   */
  Decimal: { input: number; output: number; }
};

export type Money = {
  __typename?: 'Money';
  amount: Scalars['Decimal']['output'];
  currency: Scalars['String']['output'];
};

export type MoneyInput = {
  amount: Scalars['Decimal']['input'];
  currency: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createTransaction: Transaction;
};


export type MutationCreateTransactionArgs = {
  input: NewTransaction;
};

export type NewTransaction = {
  amount: Scalars['Int']['input'];
  broker: Scalars['String']['input'];
  /**
   * If the broker does not supply an ID for the transaction, an ID is generated based on
   * the rest of the data.
   */
  brokerId?: InputMaybe<Scalars['String']['input']>;
  date: Scalars['DateTime']['input'];
  exchange: Scalars['String']['input'];
  isin: Scalars['String']['input'];
  price: MoneyInput;
  transactionCost?: InputMaybe<MoneyInput>;
};

export type Query = {
  __typename?: 'Query';
  securities: Array<Security>;
};

export type Security = {
  __typename?: 'Security';
  amount: Scalars['Int']['output'];
  averagePrice: Money;
  broker: Scalars['String']['output'];
  exchange: Scalars['String']['output'];
  isin: Scalars['String']['output'];
};

export type Transaction = {
  __typename?: 'Transaction';
  amount: Scalars['Int']['output'];
  broker: Scalars['String']['output'];
  brokerId: Scalars['String']['output'];
  date: Scalars['DateTime']['output'];
  exchange: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  isin: Scalars['String']['output'];
  price: Money;
  transactionCost?: Maybe<Money>;
};

export type CurrentlyHeldSecuritiesQueryQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentlyHeldSecuritiesQueryQuery = (
  { __typename?: 'Query' }
  & { ' $fragmentRefs'?: { 'SecurityItems_QueryFragmentFragment': SecurityItems_QueryFragmentFragment } }
);

export type SecurityItems_QueryFragmentFragment = { __typename?: 'Query', securities: Array<(
    { __typename?: 'Security', isin: string }
    & { ' $fragmentRefs'?: { 'SecurityItem_SecurityFragmentFragment': SecurityItem_SecurityFragmentFragment } }
  )> } & { ' $fragmentName'?: 'SecurityItems_QueryFragmentFragment' };

export type SecurityItem_SecurityFragmentFragment = { __typename?: 'Security', isin: string, amount: number, averagePrice: (
    { __typename?: 'Money' }
    & { ' $fragmentRefs'?: { 'MoneyItem_MoneyFragmentFragment': MoneyItem_MoneyFragmentFragment } }
  ) } & { ' $fragmentName'?: 'SecurityItem_SecurityFragmentFragment' };

export type MoneyItem_MoneyFragmentFragment = { __typename?: 'Money', amount: number, currency: string } & { ' $fragmentName'?: 'MoneyItem_MoneyFragmentFragment' };

export const MoneyItem_MoneyFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"MoneyItem_MoneyFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Money"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}}]}}]} as unknown as DocumentNode<MoneyItem_MoneyFragmentFragment, unknown>;
export const SecurityItem_SecurityFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"SecurityItem_SecurityFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Security"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"isin"}},{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"averagePrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"MoneyItem_MoneyFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"MoneyItem_MoneyFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Money"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}}]}}]} as unknown as DocumentNode<SecurityItem_SecurityFragmentFragment, unknown>;
export const SecurityItems_QueryFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"SecurityItems_QueryFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Query"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"securities"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"isin"}},{"kind":"FragmentSpread","name":{"kind":"Name","value":"SecurityItem_SecurityFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"MoneyItem_MoneyFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Money"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"SecurityItem_SecurityFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Security"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"isin"}},{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"averagePrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"MoneyItem_MoneyFragment"}}]}}]}}]} as unknown as DocumentNode<SecurityItems_QueryFragmentFragment, unknown>;
export const CurrentlyHeldSecuritiesQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentlyHeldSecuritiesQuery"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"SecurityItems_QueryFragment"}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"MoneyItem_MoneyFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Money"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"currency"}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"SecurityItem_SecurityFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Security"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"isin"}},{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"averagePrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"MoneyItem_MoneyFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"SecurityItems_QueryFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Query"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"securities"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"isin"}},{"kind":"FragmentSpread","name":{"kind":"Name","value":"SecurityItem_SecurityFragment"}}]}}]}}]} as unknown as DocumentNode<CurrentlyHeldSecuritiesQueryQuery, CurrentlyHeldSecuritiesQueryQueryVariables>;
