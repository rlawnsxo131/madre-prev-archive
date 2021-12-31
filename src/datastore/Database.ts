import {
  ConnectionManager,
  getConnectionManager,
  Connection,
  createConnection,
  ConnectionOptions,
  useContainer,
} from 'typeorm';
import { Container } from 'typeorm-typedi-extensions';

/**
 * don't know if you want to manage a connection with a different name yet.
 * should consider editing it later.
 */
export default class Database {
  private readonly connectionManager: ConnectionManager;
  private readonly connectionOptions: ConnectionOptions;

  constructor() {
    useContainer(Container);
    this.connectionManager = getConnectionManager();
    this.connectionOptions = {
      name: 'default',
      type: 'mysql',
      host: process.env.TYPEORM_HOST,
      port: Number(process.env.TYPEORM_PORT),
      username: process.env.TYPEORM_USERNAME,
      password: process.env.TYPEORM_PASSWORD,
      database: process.env.TYPEORM_DATABASE,
      charset: 'utf8mb4_unicode_ci',
      connectTimeout: 10000,
      logging: 'all',
      timezone: 'Z',
      extra: {
        connectionLimit: 10,
      },
      synchronize: false,
      entities: ['src/domain/**/entity/*.entity.ts'],
    };
  }

  connect() {
    return createConnection(this.connectionOptions);
  }

  async getConnection(): Promise<Connection> {
    const CONNECTION_NAME = 'default';
    if (this.connectionManager.has(CONNECTION_NAME)) {
      const connection = this.connectionManager.get(CONNECTION_NAME);
      try {
        if (connection.isConnected) {
          await connection.close();
        }
      } catch (e) {
        throw new Error(`connection close error: ${e}`);
      }
      return connection.connect();
    }

    return this.connect();
  }
}
