import { IResolvers } from '@graphql-tools/utils';
import { userService } from '..';

interface UserArgs {
  id: number;
}

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, { id }: UserArgs) {
      const user = await userService.getUserById(id);
      return user;
    },
  },
  Mutation: {},
};

export default resolvers;
