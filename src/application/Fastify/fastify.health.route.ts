import { FastifyPluginCallback } from 'fastify';

const fastifyHealthRoute: FastifyPluginCallback = (fastify, _, done) => {
  fastify.get('/', (request, reply) => {
    const { url, ip, method, protocol } = request;
    reply.status(200).send({
      hello: 'world',
      url,
      ip,
      method,
      protocol,
    });
  });

  done();
};

export default fastifyHealthRoute;
