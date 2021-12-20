type ErrorName =
  | 'BadRequestError'
  | 'NotFoundError'
  | 'InternalServerError'
  | 'UnauthorizedError'
  | 'ForbiddenError';

interface CustomErrorParams {
  statusCode: number;
  name: ErrorName;
  message: string;
}

class CustomError extends Error {
  statusCode: number;
  name: string;
  constructor({ statusCode, name, message }: CustomErrorParams) {
    super(message);
    this.statusCode = statusCode;
    this.name = name;
  }
}

interface ThrowFastifyErrorValidationParams {
  resolver: () => boolean;
  statusCode: number;
  message: string;
  name: ErrorName;
}

function throwFastifyErrorValidation({
  resolver,
  statusCode,
  message,
  name,
}: ThrowFastifyErrorValidationParams) {
  if (!resolver()) return;
  throw new CustomError({
    statusCode,
    message,
    name,
  });
}

export default {
  throwFastifyErrorValidation,
};
