import { IResolvers } from '@graphql-tools/utils';
import { userService, userValidationService } from '..';
import { apolloErrorManager } from '../../../lib';
import { GetUserParams } from '../interface/user.interface';

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: GetUserParams) {
      userValidationService.getUserParamsValidation(id);
      const user = await userService.getUser(id);
      if (!user) {
        apolloErrorManager.throwError({
          message: 'Not Found User',
          code: 'BAD_REQUEST',
          params: { id },
        });
      }
      return user;
    },
  },
  Mutation: {},
};

export default resolvers;
