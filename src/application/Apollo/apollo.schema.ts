import { gql } from 'apollo-server-core';
import { makeExecutableSchema } from '@graphql-tools/schema';
import merge from 'lodash.merge';
import userGraphQL from '../../domain/user/graphql';
import dataGraphQL from '../../domain/data/graphql';

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
  typeDefs: [typeDef, userGraphQL.userTypeDef, dataGraphQL.dataTypeDef],
  resolvers: merge(
    resolvers,
    userGraphQL.userResolvers,
    dataGraphQL.dataResolvers,
  ),
});

export default apolloSchema;
