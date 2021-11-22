import {
  ConnectionManager,
  getConnectionManager,
  Connection,
  createConnection,
  ConnectionOptions,
} from 'typeorm';

export default class Database {
  private readonly connectionManager: ConnectionManager;
  private readonly connectionOptions: ConnectionOptions;

  constructor() {
    this.connectionManager = getConnectionManager();
    this.connectionOptions = {
      name: 'default',
      type: 'mysql',
      host: process.env.DB_HOST,
      port: Number(process.env.DB_PORT),
      username: process.env.DB_USERNAME,
      password: process.env.DB_PASSWORD,
      database: process.env.DB_DATABASE,
      charset: 'utf8mb4_unicode_ci',
      connectTimeout: 10000,
      logging: 'all',
      timezone: 'Z',
      extra: {
        connectionLimit: 10,
      },
      synchronize: false,
      entities: ['src/components/**/entity/*.entity.ts'],
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
      } catch (e) {}
      return connection.connect();
    }

    return this.connect();
  }
}
