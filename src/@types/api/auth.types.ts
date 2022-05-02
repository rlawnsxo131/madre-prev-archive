import { UserTokenProfile } from '../domain/auth.types';

export interface GetAuthResponse {
  user_token_profile: UserTokenProfile | null;
}

export interface PostAuthGoogleCheckParams {
  access_token: string;
}

export interface PostAuthGoogleCheckResponse {
  exist: boolean;
}

export interface PostAuthGoogleSignInParams {
  access_token: string;
}

export interface PostAuthGoogleSigninResponse {
  user_token_profile: UserTokenProfile;
}

export interface PostAuthGoogleSignUpParams {
  access_token: string;
  display_name: string;
}

export interface PostAuthGoogleSignUpResponse {
  user_token_profile: UserTokenProfile;
}

export interface DeleteAuthResponse {
  is_success: boolean;
}
