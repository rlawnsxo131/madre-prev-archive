export const IS_PRODUCTION = process.env.NODE_ENV === 'production';
export const ENVIRONMENT_FILENAME = `.env.${process.env.NODE_ENV}`;

/**
 * let's think a little about how to separate the constants related to logic.
 */
export const ERROR_CODE = {
  NOT_FOUND: 'NOT_FOUND',
  FORBIDDEN: 'FORBIDDEN',
  UNAUTHENTICATED: 'UNAUTHENTICATED',
  INTERNAL_SERVER_ERROR: 'INTERNAL_SERVER_ERROR',
  BAD_REQUEST: 'BAD_REQUEST',
  BAD_USER_INPUT: 'BAD_USER_INPUT',
};
