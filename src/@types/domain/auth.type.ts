export type SocialAccountProvider = 'GOOGLE';
export interface SocialAccount {
  id: number;
  uuid: string;
  user_id: number;
  provider: SocialAccountProvider;
  social_id: string;
  created_at: string;
  updated_at: string;
}

export interface UserTokenProfile {
  display_name: string;
  email: string;
  photo_url?: string;
  access_token: string;
}
