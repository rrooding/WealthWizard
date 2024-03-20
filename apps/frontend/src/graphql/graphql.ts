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
  Amount: Scalars['Int']['input'];
  Broker: Scalars['String']['input'];
  /**
   * If the broker does not supply an ID for the transaction, an ID is generated based on
   * the rest of the data.
   */
  BrokerID?: InputMaybe<Scalars['String']['input']>;
  Date: Scalars['DateTime']['input'];
  Exchange: Scalars['String']['input'];
  ISIN: Scalars['String']['input'];
  Price: MoneyInput;
  TransactionCost?: InputMaybe<MoneyInput>;
};

export type Query = {
  __typename?: 'Query';
  securities: Array<Security>;
};

export type Security = {
  __typename?: 'Security';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  symbol: Scalars['String']['output'];
};

export type Transaction = {
  __typename?: 'Transaction';
  Amount: Scalars['Int']['output'];
  Broker: Scalars['String']['output'];
  BrokerID: Scalars['String']['output'];
  Date: Scalars['DateTime']['output'];
  Exchange: Scalars['String']['output'];
  ID: Scalars['ID']['output'];
  ISIN: Scalars['String']['output'];
  Price: Money;
  TransactionCost?: Maybe<Money>;
};
