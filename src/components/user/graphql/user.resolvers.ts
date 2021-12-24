import { IResolvers } from '@graphql-tools/utils';
import { userService, userValidationService } from '..';
import { apolloErrorService } from '../../common';
import { GetUserParams } from '../interface/user.interface';

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: GetUserParams) {
      userValidationService.getUserParamsValidation(id);
      const user = await userService.getUser(id);
      if (!user) {
        apolloErrorService.throwError({
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
