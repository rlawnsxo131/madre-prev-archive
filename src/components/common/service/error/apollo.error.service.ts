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
  | 'NOT_FOUND' // my custom
  | 'BAD_REQUEST'; // my custom

interface ThrowApolloErrorValidationParams {
  resolver: () => boolean;
  message: string;
  code: ApolloErrorCode;
  params?: Record<string, any>;
}

function throwApolloErrorValidation({
  resolver,
  message,
  code,
  params = {},
}: ThrowApolloErrorValidationParams) {
  if (!resolver()) return;
  throw new ApolloError(message, code, params);
}

export default {
  throwApolloErrorValidation,
};
