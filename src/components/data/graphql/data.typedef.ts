import { gql } from 'apollo-server-core';

export default gql`
  type Data {
    id: Int!
    data: [Int]
  }
  extend type Query {
    data(id: Int!): Data
  }
`;
