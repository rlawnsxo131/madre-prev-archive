export function isExistMapValidate(
  map: Map<any, any>,
  message: string = 'isExistMapValidate error',
) {
  if (map.size) return;
  throw new Error(message);
}
