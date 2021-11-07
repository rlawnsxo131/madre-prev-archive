import { ApolloError } from 'apollo-server-core';
import { errorCode } from '.';

interface CreateErrorParams {
  message: string;
  errorCode: errorCode;
  params: Record<string, any>;
}

function createApolloError({ message, errorCode, params }: CreateErrorParams) {
  return new ApolloError(message, errorCode, params);
}

export default {
  createApolloError,
};
