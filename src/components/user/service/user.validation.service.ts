function getUserParamsValidation(id: string) {
  if (!id) return false;
  return true;
}

export default {
  getUserParamsValidation,
};
