export function isExistMapValidate(
  map: Map<any, any>,
  message = 'isExistMapValidate error',
) {
  if (map.size) return;
  throw new Error(message);
}
