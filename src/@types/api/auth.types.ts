import { UserProfile } from '../domain/auth.types';

// get user
export type GetAuthResponse = UserProfile | null;

// check google
export interface PostAuthGoogleCheckParams {
  access_token: string;
}
export interface PostAuthGoogleCheckResponse {
  exist: boolean;
}

// sign in
export interface PostAuthGoogleSignInParams {
  access_token: string;
}
export type PostAuthGoogleSigninResponse = UserProfile;

// sign up
export interface PostAuthGoogleSignUpParams {
  access_token: string;
  username: string;
}
export type PostAuthGoogleSignUpResponse = UserProfile;
