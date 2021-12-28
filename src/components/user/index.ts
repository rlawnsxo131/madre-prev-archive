import UserRepository from './repository/user.repository';
import UserQueryRepository from './repository/user.query.repository';
import userResolvers from './graphql/user.resolvers';
import userTypeDef from './graphql/user.typedef';
import userService from './service/user.service';
import userValidationService from './service/user.validation.service';

const userGraphQL = {
  userTypeDef,
  userResolvers,
};

export {
  UserRepository,
  UserQueryRepository,
  userService,
  userValidationService,
  userGraphQL,
};
