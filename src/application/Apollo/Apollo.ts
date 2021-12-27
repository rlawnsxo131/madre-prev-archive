import { ApolloServer } from 'apollo-server-fastify';
import {
  ApolloServerPluginDrainHttpServer,
  ApolloServerPluginLandingPageDisabled,
  ApolloServerPluginLandingPageGraphQLPlayground,
} from 'apollo-server-core';
import { FastifyInstance } from 'fastify';
import apolloSchema from './apollo.schema';
import { isProduction } from '../../constants';

export default class Apollo {
  private readonly app: ApolloServer;

  constructor(fastify: FastifyInstance) {
    this.app = new ApolloServer({
      schema: apolloSchema,
      context: ({ request }) => {
        request.log.info(request.id);
      },
      plugins: [
        this.fastifyAppClosePlugin(fastify),
        ApolloServerPluginDrainHttpServer({ httpServer: fastify.server }),
        isProduction
          ? ApolloServerPluginLandingPageDisabled()
          : ApolloServerPluginLandingPageGraphQLPlayground(),
      ],
      debug: !isProduction,
    });
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
