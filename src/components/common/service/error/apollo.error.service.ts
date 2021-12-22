import { ApolloError } from 'apollo-server-core';

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
  | 'BAD_REQUEST'; // my custom

interface ThrowErrorValidationParams {
  resolver: (params?: any) => boolean;
  message: string;
  code: ErrorCode;
  params?: Record<string, any>;
}

function throwErrorValidation({
  resolver,
  message,
  code,
  params = {},
}: ThrowErrorValidationParams) {
  if (!resolver()) return;
  throw new ApolloError(message, code, params);
}

export default {
  throwErrorValidation,
};
