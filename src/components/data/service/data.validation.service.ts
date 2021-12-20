import { CreateDataParams } from '../interface/data.interface';

function getDataParamsValidation(id: string) {
  if (!id) return false;
  return true;
}

function createDataParamsValidation(params: CreateDataParams) {
  if (!params.user_id) return false;
  if (!params.file_url) return false;
  if (!params.title) return false;
  return true;
}

export default {
  getDataParamsValidation,
  createDataParamsValidation,
};
