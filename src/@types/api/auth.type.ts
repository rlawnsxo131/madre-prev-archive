export interface PostAuthGoogleSigninParams {
  accessToken: string;
}

export interface PostAuthGoogleSigninResponse {
  exist: boolean;
}

export interface PostAuthGoogleSignupParams {
  accessToken: string;
  username: string;
}

export interface PostAuthGoogleSignupResponse {}
