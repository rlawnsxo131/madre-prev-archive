/**
 * let's think a little more
 */
interface CustomErrorParams {
  statusCode: number;
  name: string;
  message: string;
}
export default class CustomError extends Error {
  statusCode: number;
  name: string;
  constructor({ statusCode, name, message }: CustomErrorParams) {
    super(message);
    this.statusCode = statusCode;
    this.name = name;
  }
}
