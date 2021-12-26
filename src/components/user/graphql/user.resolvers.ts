import { IResolvers } from '@graphql-tools/utils';
import { UserService, UserValidationService } from '..';
import { ApolloErrorUtil } from '../../../utils';
import { GetUserParams } from '../interface/user.interface';

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: GetUserParams) {
      UserValidationService.getUserParamsValidation(id);
      const user = await UserService.getUser(id);
      if (!user) {
        ApolloErrorUtil.throwError({
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
