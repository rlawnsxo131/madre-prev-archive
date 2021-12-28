export interface GetDataParams {
  id: string;
}

export interface CreateDataParams {
  user_id: string;
  file_url: string;
  title: string;
  description?: string;
  is_public: boolean;
}
