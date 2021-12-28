import DataQueryRepository from './repository/data.query.repository';
import DataRepository from './repository/data.repository';
import dataService from './service/data.service';
import dataValidationService from './service/data.validation.service';
import dataTypeDef from './graphql/data.typedef';
import dataResolvers from './graphql/data.resolvers';

const dataGraphQL = {
  dataTypeDef,
  dataResolvers,
};

export {
  DataRepository,
  DataQueryRepository,
  dataService,
  dataValidationService,
  dataGraphQL,
};
