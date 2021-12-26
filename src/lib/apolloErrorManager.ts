import { ApolloError } from 'apollo-server-core';

interface ThrowApolloErrorParams {
  message: string;
  code:
    | 'GRAPHQL_PARSE_FAILED'
    | 'GRAPHQL_VALIDATION_FAILED'
    | 'BAD_USER_INPUT'
    | 'UNAUTHENTICATED'
    | 'FORBIDDEN'
    | 'PERSISTED_QUERY_NOT_FOUND'
    | 'PERSISTED_QUERY_NOT_SUPPORTED'
    | 'INTERNAL_SERVER_ERROR'
    | 'NOT_FOUND' // my custom
    | 'BAD_REQUEST'; // my custom;;
  params?: Record<string, any>;
}

function throwError({ message, code, params = {} }: ThrowApolloErrorParams) {
  throw new ApolloError(message, code, params);
}

export default {
  throwError,
};
