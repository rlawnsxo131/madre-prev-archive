export type SocialAccountProvider = 'GOOGLE';

export interface Account {
  user_id: number;
  username: string;
  email: string;
  photo_url: string;
  social_id: string;
  provider: SocialAccountProvider;
  created_at: string;
  updated_at: string;
}

export interface SocialAccount {
  id: string;
  user_id: string;
  provider: SocialAccountProvider;
  social_id: string;
  created_at: string;
  updated_at: string;
}

export interface UserProfile {
  user_id: string;
  username: string;
  photo_url?: string;
}
