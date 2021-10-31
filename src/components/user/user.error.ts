import { ForbiddenError } from 'apollo-server-core';

const notFoundUser = new ForbiddenError('Not Found User');

export default {
  notFoundUser,
};
