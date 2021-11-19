import { gql } from 'apollo-server-core';

export default gql`
  type User {
    id: Int!
    name: String!
  }
  extend type Query {
    user(id: Int!): User
  }
`;
