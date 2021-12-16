import { IResolvers } from '@graphql-tools/utils';
import { userService } from '..';
import { errorService } from '../../error';

interface UserArgs {
  id: number;
}

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: UserArgs) {
      const user = await userService.getUser(id);
      errorService.throwApolloError({
        resolver: () => !user,
        message: 'Not Found Data',
        code: 'NOT_FOUND',
        params: { id },
      });
      return user;
    },
  },
  Mutation: {},
};

export default resolvers;
