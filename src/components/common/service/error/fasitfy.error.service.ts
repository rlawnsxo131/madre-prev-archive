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

interface ThrowErrorValidationParams {
  resolver: (params?: any) => boolean;
  statusCode: number;
  message: string;
  name: ErrorName;
}

function throwErrorValidation({
  resolver,
  statusCode,
  message,
  name,
}: ThrowErrorValidationParams) {
  if (!resolver()) return;
  throw new CustomError({
    statusCode,
    message,
    name,
  });
}

export default {
  throwErrorValidation,
};
