import { atom, useRecoilValue, useSetRecoilState } from 'recoil';

interface UserState {}

const userState = atom<UserState>({
  key: 'userState',
  default: {},
});

export function useUserValue() {
  return useRecoilValue(userState);
}

export function useUserActions() {
  const set = useSetRecoilState(userState);

  return {};
}
