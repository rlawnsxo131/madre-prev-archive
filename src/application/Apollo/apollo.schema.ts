import { gql } from 'apollo-server-core';
import { makeExecutableSchema } from '@graphql-tools/schema';
import merge from 'lodash.merge';
import UserGraphQL from '../../domain/user/graphql';
import DataGraphQL from '../../domain/data/graphql';

const typeDef = gql`
  scalar Date
  type Query {
    _version: String
  }
  type Mutation {
    _empty: String
  }
`;

const resolvers = {
  Query: {
    _version: () => '1.0',
  },
  Mutation: {},
};

const apolloSchema = makeExecutableSchema({
  typeDefs: [typeDef, UserGraphQL.userTypeDef, DataGraphQL.dataTypeDef],
  resolvers: merge(
    resolvers,
    UserGraphQL.userResolvers,
    DataGraphQL.dataResolvers,
  ),
});

export default apolloSchema;
