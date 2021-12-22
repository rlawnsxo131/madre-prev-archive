import { IResolvers } from '@graphql-tools/utils';
import { dataService, dataValidationService } from '..';
import { apolloErrorService } from '../../common';
import { CreateDataParams, GetDataParams } from '../interface/data.interface';

const resolvers: IResolvers = {
  Query: {
    async data(_, { id }: GetDataParams) {
      dataValidationService.getDataParamsValidation(id);
      const data = await dataService.getData(id);
      apolloErrorService.throwErrorValidation({
        resolver: () => !data,
        message: 'Not Found Data',
        code: 'NOT_FOUND',
        params: { id },
      });
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
