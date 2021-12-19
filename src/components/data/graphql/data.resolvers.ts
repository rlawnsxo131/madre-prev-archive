import { IResolvers } from '@graphql-tools/utils';
import { dataService } from '..';
import { errorService } from '../../error';

interface DataArgs {
  id: string;
}

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: DataArgs) {
      const data = await dataService.getData(id);
      errorService.throwApolloError({
        resolver: () => !data,
        message: 'Not Found Data',
        code: 'NOT_FOUND',
        params: { id },
      });
      return data;
    },
  },
  Mutation: {
    async createData() {},
  },
};

export default resolvers;
