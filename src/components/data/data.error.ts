import { ForbiddenError } from 'apollo-server-core';

const notFoundData = new ForbiddenError('Not Found Data');

export default {
  notFoundData,
};
