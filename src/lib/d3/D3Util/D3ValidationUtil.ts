export function isExistMapValidate(
  map: Map<any, any>,
  message: string = 'unknown error',
) {
  if (map.size) return;
  throw new Error(message);
}
