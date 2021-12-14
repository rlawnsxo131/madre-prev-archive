import { IResolvers } from '@graphql-tools/utils';
import { ApolloError } from 'apollo-server-core';
import { userService } from '..';
import { ERROR_CODE } from '../../../constants';

interface UserArgs {
  id: number;
}

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: UserArgs) {
      const user = await userService.getUser(id);
      if (!user) {
        throw new ApolloError('Not Found User', ERROR_CODE.BAD_REQUEST, { id });
      }
      return user;
    },
  },
  Mutation: {},
};

export default resolvers;
