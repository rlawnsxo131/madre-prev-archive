/**
 * let's think a little more
 */
export default class CustomError extends Error {
  statusCode: number;
  name: string;
  constructor(statusCode: number, name: string, message: string) {
    super(message);
    this.statusCode = statusCode;
    this.name = name;
  }
}
