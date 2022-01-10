import { IResolvers } from '@graphql-tools/utils';
import Joi from 'joi';
import { dataService } from '..';
import ApolloCustomError from '../../../lib/ApolloCustomError';
import apolloValidator from '../../../lib/validation/apolloValidator';
import { GetDataParams, CreateDataParams } from '../interface/data.interface';

const dataResolvers: IResolvers = {
  Query: {
    async data(_, { id }: GetDataParams) {
      apolloValidator.validateId(id);
      const data = await dataService.findOne(id);
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
      apolloValidator.validateObject(
        Joi.object<CreateDataParams>({
          user_id: Joi.string().guid().required(),
          file_url: Joi.string().uri().required(),
          title: Joi.string().min(1).required(),
          is_public: Joi.boolean().required(),
          description: Joi.string().min(1).optional(),
        }),
        params,
      );
      const data = await dataService.createOne(params);
      return data;
    },
  },
};

export default dataResolvers;
