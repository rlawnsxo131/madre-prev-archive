import { IResolvers } from '@graphql-tools/utils';
import { userActionService } from '..';

const resolvers: IResolvers = {
  User: {},
  Query: {
    async user(_, args) {
      const { id } = args;
      const user = await userActionService.getUserAction(id);
      return user;
    },
  },
  Mutation: {},
};

export default resolvers;
