export interface GetAuthResponse {
  profile: {
    display_name: string;
    email: string;
    photo_url?: string;
    access_token: string;
  };
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
  access_token: string;
  display_name: string;
  photo_url?: string;
}

export interface PostAuthGoogleSignUpParams {
  access_token: string;
  display_name: string;
}

export interface PostAuthGoogleSignUpResponse {
  access_token: string;
  display_name: string;
  photo_url?: string;
}
