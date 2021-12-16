import { IResolvers } from '@graphql-tools/utils';
import { dataService } from '..';
import { errorService } from '../../error';

interface DataArgs {
  id: number;
}

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: DataArgs) {
      const data = await dataService.getData(id);
      errorService.throwApolloError({
        resolver: () => !data,
        code: 'NOT_FOUND',
        message: 'Not Found Data',
        params: { id },
      });
      return data;
    },
  },
};

export default resolvers;
