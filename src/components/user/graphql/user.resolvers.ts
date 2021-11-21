import { IResolvers } from '@graphql-tools/utils';
import { userActionService } from '..';

interface UserArgs {
  id: number;
}

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, args: UserArgs) {
      const { id } = args;
      const user = await userActionService.getUserAction(id);
      return user;
    },
  },
  Mutation: {},
};

export default resolvers;
