import { IResolvers } from '@graphql-tools/utils';
import { dataService } from '..';
import { errorService } from '../../error';
import { CreateDataParams, GetDataParams } from '../interface/data.interface';

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: GetDataParams) {
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
    async createData(_, args: CreateDataParams) {
      const data = await dataService.createData(args);
      return data;
    },
  },
};

export default resolvers;
