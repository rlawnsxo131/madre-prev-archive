import { IResolvers } from '@graphql-tools/utils';
import { dataActionService } from '..';

const resolvers: IResolvers = {
  Query: {
    async data(_, args) {
      const { id } = args;
      const data = await dataActionService.getDataAction(id);
      return data;
    },
  },
};

export default resolvers;
