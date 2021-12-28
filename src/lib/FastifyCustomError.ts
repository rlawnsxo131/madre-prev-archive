type FastifyErrorName =
  | 'BadRequestError'
  | 'NotFoundError'
  | 'InternalServerError'
  | 'UnauthorizedError'
  | 'ForbiddenError';

interface FastifyCustomErrorParams {
  message: string;
  name: FastifyErrorName;
  statusCode: number;
}

export default class FastifyCustomError extends Error {
  name: FastifyErrorName;
  statusCode: number;
  constructor({ message, name, statusCode }: FastifyCustomErrorParams) {
    super(message);
    this.name = name;
    this.statusCode = statusCode;
  }
}
