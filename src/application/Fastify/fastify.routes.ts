import { FastifyPluginCallback } from 'fastify';
import { authRoute } from '../../components/auth';

const routes: FastifyPluginCallback = (fastify, opts, done) => {
  fastify.register(authRoute, { prefix: '/auth' });

  fastify.get('/health', (request, reply) => {
    reply.status(200).send({ hello: 'world' });
  });

  done();
};

export default routes;
