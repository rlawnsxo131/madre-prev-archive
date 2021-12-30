function isExistMapValidate(
  map: Map<any, any>,
  message: string = 'unknown error',
) {
  if (map.size) return;
  throw new Error(message);
}

const D3ValidationUtil = {
  isExistMapValidate,
};

export default D3ValidationUtil;
