import { ApolloError } from 'apollo-server-core';
import { errorCode } from '.';

interface CreateApolloErrorParams {
  message: string;
  errorCode: errorCode;
  params: Record<string, any>;
}

function createApolloError({
  message,
  errorCode,
  params,
}: CreateApolloErrorParams) {
  return new ApolloError(message, errorCode, params);
}

export default {
  createApolloError,
};
