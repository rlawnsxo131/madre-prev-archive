import useScreenSignupActions from '../../../../hooks/screenSignup/useScreenSignupActions';
import useScreenSignupState from '../../../../hooks/screenSignup/useScreenSignupState';
import useInputs from '../../../../hooks/useInputs';
import { normalizedString } from '../../../../lib/utils';
import authApi from '../../../../store/api/authApi';

export default function useScreenSignupRightBlock() {
  const { accessToken } = useScreenSignupState();
  const [googleSignup] = authApi.usePostGoogleSignupMutation();
  const { close } = useScreenSignupActions();
  const { state, onChange } = useInputs({
    username: '',
  });

  const onSignup = async () => {
    const normalAccessToken = normalizedString(accessToken);
    const normalUsername = normalizedString(state.username);

    if (!normalAccessToken || !normalUsername) {
      if (!normalAccessToken) {
        console.log('시스템 에러');
        return;
      }
      return;
    }

    await googleSignup({
      accessToken: normalAccessToken,
      username: normalUsername,
    });
  };

  return {
    state,
    close,
    onChange,
    onSignup,
  };
}
