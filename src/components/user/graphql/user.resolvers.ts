import { IResolvers } from '@graphql-tools/utils';
import { userService } from '..';
import { apolloErrorService } from '../../common';
import { GetUserParams } from '../interface/user.interface';
import userValidationService from '../service/user.validation.service';

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: GetUserParams) {
      userValidationService.getUserParamsValidation(id);
      const user = await userService.getUser(id);
      apolloErrorService.throwApolloErrorValidation({
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
