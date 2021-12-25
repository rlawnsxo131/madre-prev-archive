import { gql } from 'apollo-server-core';
import { makeExecutableSchema } from '@graphql-tools/schema';
import merge from 'lodash.merge';
import { userGraphQL } from '../../components/user';
import { dataGraphQL } from '../../components/data';

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

const schema = makeExecutableSchema({
  typeDefs: [typeDef, userGraphQL.userTypeDef, dataGraphQL.dataTypeDef],
  resolvers: merge(
    resolvers,
    userGraphQL.userResolvers,
    dataGraphQL.dataResolvers,
  ),
});

export default schema;
