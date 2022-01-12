import fastify, {
  FastifyError,
  FastifyInstance,
  FastifyReply,
  FastifyRequest,
} from 'fastify';
import corsPlugin from 'fastify-cors';
import compress from 'fastify-compress';
import cookie from 'fastify-cookie';
import { jwtPlugin } from '../../plugins';
import fastifyRoutes from './fastify.routes';
import fastifyHealthRoute from './fastify.health.route';

export default class Fastify {
  private readonly app: FastifyInstance;

  constructor() {
    this.app = fastify({
      logger: {
        prettyPrint: {
          colorize: true,
          translateTime: 'SYS:standard',
        },
      },
    });
    this.app.setErrorHandler(this.errorHandler);
    this.app.register(corsPlugin, {
      origin: (origin, callback) => {
        const allowedHost = [/^http\:\/\/localhost/];
        const allowed = allowedHost.some((regex) => regex.test(origin));
        callback(null, allowed);
      },
      credentials: true,
    });
    this.app.register(cookie);
    this.app.register(compress);
    this.app.register(jwtPlugin);
    this.app.register(fastifyHealthRoute, { prefix: '/health' });
    this.app.register(fastifyRoutes, { prefix: '/api' });
  }

  private errorHandler(
    error: FastifyError,
    _: FastifyRequest,
    reply: FastifyReply,
  ) {
    const statusCode = error.statusCode || 500;
    reply.status(statusCode);
    reply.send({
      statusCode: statusCode,
      name: error.name,
      message: error.message,
    });
  }

  getApp() {
    return this.app;
  }

  start() {
    return this.app.listen(process.env.PORT!);
  }

  close() {
    return this.app.close();
  }

  registerApollo(apolloHandler: (fastify: FastifyInstance) => Promise<void>) {
    this.app.register(apolloHandler);
  }
}
