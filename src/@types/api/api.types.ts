/**
 * rtk catch block error response data set
 * {
 *   error: {
 *     status: 500,
 *     data: {
 *       status: 500
 *       code: "InternalServerError"
 *     }
 *   }
 *   isUnhandledError: false
 *   meta: {request: Request, response: Response}
 * }
 */

type ResponseErrors =
  | {
      status: 400;
      code: 'BadRequest';
    }
  | {
      status: 401;
      code: 'Unauthorized';
    }
  | {
      status: 403;
      code: 'Forbidden';
    }
  | {
      status: 404;
      code: 'NotFound';
    }
  | {
      status: 409;
      code: 'Conflict';
    }
  | {
      status: 422;
      code: 'UnprocessableEntity';
    }
  | {
      status: 500;
      code: 'InternalServerError';
    };

export interface ResponseError {
  error: {
    status: number;
    data: ResponseErrors;
  };
  isUnhandledError: boolean;
  meta: {
    request: Request;
    response: Response;
  };
}
