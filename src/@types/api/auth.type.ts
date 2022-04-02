export interface PostAuthGoogleSignInParams {
  access_token: string;
}

export interface PostAuthGoogleSigInResponse {
  exist: boolean;
}

export interface PostAuthGoogleSignUpParams {
  access_token: string;
  display_name: string;
}

export interface PostAuthGoogleSignUpResponse {}
