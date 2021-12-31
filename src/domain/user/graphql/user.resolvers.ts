import { IResolvers } from '@graphql-tools/utils';
import { UserService } from '..';
import ApolloCustomError from '../../../lib/ApolloCustomError';
import { GetUserParams } from '../interface/user.interface';
import { ApolloValidator } from '../../../lib/Validator';
import Container from 'typedi';

const userResolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: GetUserParams) {
      ApolloValidator.validateId(id);
      const user = await Container.get(UserService).findOne(id);
      if (!user) {
        throw new ApolloCustomError({
          message: 'Not Found User',
          code: 'BAD_REQUEST',
          extensions: { id },
        });
      }
      return user;
    },
  },
  Mutation: {},
};

export default userResolvers;
