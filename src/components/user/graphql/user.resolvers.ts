import { IResolvers } from '@graphql-tools/utils';
import { userService, userValidationService } from '..';
import { errorService } from '../../error';
import { GetUserParams } from '../interface/user.interface';

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: GetUserParams) {
      const validation = userValidationService.getUserParamsValidation(id);
      errorService.throwApolloError({
        resolver: () => !validation,
        message: 'user query validation error',
        code: 'BAD_USER_INPUT',
        params: { id },
      });
      const user = await userService.getUser(id);
      errorService.throwApolloError({
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
