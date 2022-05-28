// object keys
export const MADRE_COLOR_THEME = 'MADRE_COLOR_THEME';
export const MADRE_USER_PROFILE = 'MADRE_USER_PROFILE';

// variables
export const isProduction = process.env.REACT_APP_NODE_ENV === 'production';

export const userPath = '/@:username';
export const appRoutes = [
  { path: '/' },
  { path: '/madre-story' },
  { path: '/notice' },
  { path: '/guide' },
  { path: '/policy' },
  { path: '/notifications' },
  { path: userPath },
  { path: '/m/all-menu' },
];

// The values below are in the order in which they are displayed.
export const appInfoRoutes = [
  { path: '/madre-story', displayText: 'Madre 이야기' },
  { path: '/notice', displayText: '공지사항' },
  { path: '/guide', displayText: '가이드 및 튜토리얼' },
];
