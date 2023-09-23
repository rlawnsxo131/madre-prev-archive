import { ApolloError } from 'apollo-server-core';

type ApolloErrorCode =
  | 'GRAPHQL_PARSE_FAILED'
  | 'GRAPHQL_VALIDATION_FAILED'
  | 'BAD_USER_INPUT'
  | 'UNAUTHENTICATED'
  | 'FORBIDDEN'
  | 'PERSISTED_QUERY_NOT_FOUND'
  | 'PERSISTED_QUERY_NOT_SUPPORTED'
  | 'INTERNAL_SERVER_ERROR'
  /**
   * custom error code
   */
  | 'NOT_FOUND'
  | 'BAD_REQUEST';

interface ApolloCustomErrorParams {
  message: string;
  code: ApolloErrorCode;
  extensions?: Record<string, any>;
}

export default class ApolloCustomError extends ApolloError {
  constructor({ message, code, extensions = {} }: ApolloCustomErrorParams) {
    super(message, code, extensions);
  }
}
