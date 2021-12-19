import { IResolvers } from '@graphql-tools/utils';
import { userService } from '..';
import { errorService } from '../../error';

interface UserArgs {
  id: string;
}

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: UserArgs) {
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
