// object keys
export const MADRE_COLOR_THEME = 'MADRE_COLOR_THEME';
export const MADRE_USER_TOKEN_PROFILE = 'MADRE_USER_TOKEN_PROFILE';

// variables
export const isProduction = process.env.REACT_APP_NODE_ENV === 'production';

export const userPath = '/@:displayName';
export const appRoutes = [
  { path: '/' },
  { path: '/madre-story' },
  { path: '/notice' },
  { path: '/guide' },
  { path: '/policy' },
  { path: userPath },
];

// The values ​​below are in the order in which they are displayed.
export const appDisplayRoutes = [
  { path: '/madre-story', displayName: 'Madre 이야기' },
  { path: '/notice', displayName: '공지사항' },
  { path: '/guide', displayName: '가이드 및 튜토리얼' },
];
