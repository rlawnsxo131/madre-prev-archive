import { FastifyPluginCallback } from 'fastify';

const route: FastifyPluginCallback = (fastify, _, done) => {
  fastify.get('/', (_, reply) => {
    reply.status(200).send({
      auth: 'auth',
    });
  });
  done();
};

export default route;
