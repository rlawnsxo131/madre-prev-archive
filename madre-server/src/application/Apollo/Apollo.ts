import { ApolloServer } from 'apollo-server-fastify';
import {
  ApolloServerPluginDrainHttpServer,
  ApolloServerPluginLandingPageDisabled,
  ApolloServerPluginLandingPageGraphQLPlayground,
  GraphQLRequestContext,
} from 'apollo-server-core';
import { FastifyInstance } from 'fastify';
import apolloSchema from './apollo.schema';
import { isProduction } from '../../constants';
import { GraphQLError } from 'graphql';
import logger from '../../lib/logger';

export default class Apollo {
  private readonly app: ApolloServer;

  constructor(fastify: FastifyInstance) {
    this.app = new ApolloServer({
      schema: apolloSchema,
      context: ({ request, reply }) => {
        request.log.info('context');
      },
      formatError: this.formatError,
      plugins: [
        this.fastifyAppClosePlugin(fastify),
        ApolloServerPluginDrainHttpServer({ httpServer: fastify.server }),
        this.apolloRequestDidStartPlugin(),
        isProduction
          ? ApolloServerPluginLandingPageDisabled()
          : ApolloServerPluginLandingPageGraphQLPlayground(),
      ],
      debug: !isProduction,
    });
  }

  private formatError(error: GraphQLError) {
    logger.error(`${error}\n${error.extensions.exception?.stacktrace}`);
    return error;
  }

  private fastifyAppClosePlugin(fastify: FastifyInstance) {
    return {
      async serverWillStart() {
        return {
          async drainServer() {
            await fastify.close();
          },
        };
      },
    };
  }

  private apolloRequestDidStartPlugin() {
    return {
      async requestDidStart(requestContext: GraphQLRequestContext) {
        logger.info('GraphQL requestDidStart');
        logger.info(`${requestContext.request.query}`);
      },
    };
  }

  getApp() {
    return this.app;
  }

  start() {
    return this.app.start();
  }

  createHandler() {
    return this.app.createHandler({ cors: false });
  }
}
