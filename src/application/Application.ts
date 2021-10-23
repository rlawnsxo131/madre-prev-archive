import Fastify from './Fastify';
import Apollo from './Apollo';

export default class Application {
  private fastify: Fastify;
  private apollo: Apollo;

  constructor() {
    this.fastify = new Fastify();
    this.apollo = new Apollo(this.fastify.getApp());
  }

  getFastify() {
    return this.fastify;
  }

  getApollo() {
    return this.apollo;
  }

  async setup() {
    await this.apollo.start();
    this.fastify.registerApollo(this.apollo.createHandler());
  }

  async start() {
    try {
      await this.fastify.start();
    } catch (e) {
      this.fastify.getApp().log.error(e);
      process.exit(1);
    }
  }
}
