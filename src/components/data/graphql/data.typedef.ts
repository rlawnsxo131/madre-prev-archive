import { gql } from 'apollo-server-core';

const dataTypeDef = gql`
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
    data(id: ID!): Data!
  }

  extend type Mutation {
    createData(
      user_id: String!
      file_url: String!
      title: String!
      description: String
      is_public: Boolean!
    ): Data!
  }
`;

export default dataTypeDef;
