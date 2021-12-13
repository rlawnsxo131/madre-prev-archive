import { IResolvers } from '@graphql-tools/utils';
import { dataService } from '..';

interface DataArgs {
  id: number;
}

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: DataArgs) {
      const data = await dataService.getDataById(id);
      return data;
    },
  },
};

export default resolvers;
