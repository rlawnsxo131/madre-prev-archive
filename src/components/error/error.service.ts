import { ApolloError } from 'apollo-server-core';

const ERROR_CODE = {
  NOT_FOUND: 'NOT_FOUND',
  FORBIDDEN: 'FORBIDDEN',
  UNAUTHENTICATED: 'UNAUTHENTICATED',
  INTERNAL_SERVER_ERROR: 'INTERNAL_SERVER_ERROR',
  BAD_REQUEST: 'BAD_REQUEST',
  BAD_USER_INPUT: 'BAD_USER_INPUT',
};

interface ThrowApolloErrorParams {
  resolver: () => boolean;
  message: string;
  code: keyof typeof ERROR_CODE;
  params?: any;
}

function throwApolloError({
  resolver,
  message,
  code,
  params = {},
}: ThrowApolloErrorParams) {
  if (!resolver()) return;
  throw new ApolloError(message, ERROR_CODE[code], { ...params });
}

export default {
  ERROR_CODE,
  throwApolloError,
};
