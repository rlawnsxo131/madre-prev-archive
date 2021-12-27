import { IResolvers } from '@graphql-tools/utils';
import { dataService, dataValidationService } from '..';
import ApolloCustomError from '../../../lib/ApolloCustomError';
import { CreateDataParams, GetDataParams } from '../interface/data.interface';

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: GetDataParams) {
      dataValidationService.getDataParamsValidation(id);
      const data = await dataService.getData(id);
      if (!data) {
        throw new ApolloCustomError({
          message: 'Not Found Data',
          code: 'NOT_FOUND',
          extensions: { id },
        });
      }
      return data;
    },
  },
  Mutation: {
    async createData(_, params: CreateDataParams) {
      dataValidationService.createDataParamsValidation(params);
      const data = await dataService.createData(params);
      return data;
    },
  },
};

export default resolvers;
