import { gql } from 'apollo-server-core';

export default gql`
  type Data {
    id: ID!
    user_id: ID!
    file_url: String!
    title: String!
    description: String
    is_public: Boolean!
    created_at: Date!
    updated_at: Date!
  }
  extend type Query {
    data(id: Int!): Data!
  }
  extend type Mutation {
    createData: Data!
  }
`;
