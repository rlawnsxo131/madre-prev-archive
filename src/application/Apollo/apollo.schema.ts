import { gql } from 'apollo-server-core';
import { makeExecutableSchema } from '@graphql-tools/schema';
import merge from 'lodash.merge';
import { UserGraphQL } from '../../components/user';
import { DataGraphQL } from '../../components/data';

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
  typeDefs: [typeDef, UserGraphQL.typeDef, DataGraphQL.typeDef],
  resolvers: merge(resolvers, UserGraphQL.resolvers, DataGraphQL.resolvers),
});

export default schema;
