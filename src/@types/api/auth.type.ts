export interface PostAuthGoogleSignInParams {
  accessToken: string;
}

export interface PostAuthGoogleSigInResponse {
  exist: boolean;
}

export interface PostAuthGoogleSignUpParams {
  accessToken: string;
  username: string;
}

export interface PostAuthGoogleSignUpResponse {}
