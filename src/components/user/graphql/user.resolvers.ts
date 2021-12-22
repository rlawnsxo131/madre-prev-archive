import { IResolvers } from '@graphql-tools/utils';
import { userService } from '..';
import { apolloErrorService } from '../../common';
import { GetUserParams } from '../interface/user.interface';

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: GetUserParams) {
      const user = await userService.getUser(id);
      apolloErrorService.throwErrorValidation({
        resolver: () => !user,
        message: 'Not Found User',
        code: 'BAD_REQUEST',
        params: { id },
      });
      return user;
    },
  },
  Mutation: {},
};

export default resolvers;
