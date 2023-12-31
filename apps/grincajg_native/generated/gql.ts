/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

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
    "mutation createUser($input: CreateUserInput!) {\n  createUser(input: $input) {\n    name\n    id\n  }\n}": types.CreateUserDocument,
    "query getMainCategories {\n  categories {\n    id\n    name\n  }\n}": types.GetMainCategoriesDocument,
    "query getOrganizationsOnMap($input: getOrganizationsOnMapInput!) {\n  getOrganizationsOnMap(input: $input) {\n    id\n    latitude\n    longitude\n  }\n}": types.GetOrganizationsOnMapDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "mutation createUser($input: CreateUserInput!) {\n  createUser(input: $input) {\n    name\n    id\n  }\n}"): (typeof documents)["mutation createUser($input: CreateUserInput!) {\n  createUser(input: $input) {\n    name\n    id\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query getMainCategories {\n  categories {\n    id\n    name\n  }\n}"): (typeof documents)["query getMainCategories {\n  categories {\n    id\n    name\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query getOrganizationsOnMap($input: getOrganizationsOnMapInput!) {\n  getOrganizationsOnMap(input: $input) {\n    id\n    latitude\n    longitude\n  }\n}"): (typeof documents)["query getOrganizationsOnMap($input: getOrganizationsOnMapInput!) {\n  getOrganizationsOnMap(input: $input) {\n    id\n    latitude\n    longitude\n  }\n}"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;