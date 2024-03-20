/* eslint-disable */
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
  DateTime: { input: any; output: any; }
  /**
   * A signed decimal number, which supports arbitrary precision and is serialized as a string.
   *
   * Example values: 29.99, 29.999.
   */
  Decimal: { input: any; output: any; }
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
