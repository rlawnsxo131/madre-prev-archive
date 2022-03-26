export interface User {
  id: number;
  uuid: string;
  authId: number;
  email: string;
  username?: string;
  display_name: string;
  photo_url?: string;
  created_at: string;
  updated_at: string;
}
