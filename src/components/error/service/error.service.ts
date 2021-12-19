import { ApolloError } from 'apollo-server-core';

interface ThrowApolloErrorParams {
  resolver: () => boolean;
  message: string;
  code: ErrorCode;
  params?: Record<string, any>;
}

function throwApolloError({
  resolver,
  message,
  code,
  params = {},
}: ThrowApolloErrorParams) {
  if (!resolver()) return;
  throw new ApolloError(message, code, { ...params });
}

export default {
  throwApolloError,
};
