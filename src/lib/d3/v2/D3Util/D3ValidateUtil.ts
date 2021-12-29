function isExistMapValidation(
  map: Map<any, any>,
  message: string = 'unknown error',
) {
  if (map.size) return;
  throw new Error(message);
}

const D3ValidateUtil = {
  isExistMapValidation,
};

export default D3ValidateUtil;
