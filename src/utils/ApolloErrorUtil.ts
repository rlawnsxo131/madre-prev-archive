import { ApolloError } from 'apollo-server-core';

export namespace ApolloErrorUtil {
  type ErrorCode =
    | 'GRAPHQL_PARSE_FAILED'
    | 'GRAPHQL_VALIDATION_FAILED'
    | 'BAD_USER_INPUT'
    | 'UNAUTHENTICATED'
    | 'FORBIDDEN'
    | 'PERSISTED_QUERY_NOT_FOUND'
    | 'PERSISTED_QUERY_NOT_SUPPORTED'
    | 'INTERNAL_SERVER_ERROR'
    | 'NOT_FOUND' // my custom
    | 'BAD_REQUEST'; // my custom;

  interface ThrowErrorParams {
    message: string;
    code: ErrorCode;
    params?: Record<string, any>;
  }

  export function throwError({ message, code, params = {} }: ThrowErrorParams) {
    throw new ApolloError(message, code, params);
  }
}
