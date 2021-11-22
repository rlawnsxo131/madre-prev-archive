import { IResolvers } from '@graphql-tools/utils';
import { dataActionService } from '..';

interface DataArgs {
  id: number;
}

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: DataArgs) {
      const data = await dataActionService.getDataAction(id);
      return data;
    },
  },
};

export default resolvers;
