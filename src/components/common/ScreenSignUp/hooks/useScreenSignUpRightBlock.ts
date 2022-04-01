import useScreenSignUpActions from '../../../../hooks/screenSignUp/useScreenSignUpActions';
import useScreenSignUpState from '../../../../hooks/screenSignUp/useScreenSignUpState';
import useInputs from '../../../../hooks/useInputs';
import { normalizedString } from '../../../../lib/utils';
import authApi from '../../../../store/api/authApi';

export default function useScreenSignUpRightBlock() {
  const { accessToken } = useScreenSignUpState();
  const [googleSignUp] = authApi.usePostGoogleSignUpMutation();
  const { close } = useScreenSignUpActions();
  const { state, onChange } = useInputs({
    username: '',
  });

  const onSignUp = async () => {
    const normalAccessToken = normalizedString(accessToken);
    const normalUsername = normalizedString(state.username);

    if (!normalAccessToken || !normalUsername) {
      if (!normalAccessToken) {
        console.log('시스템 에러');
        return;
      }
      return;
    }

    await googleSignUp({
      accessToken: normalAccessToken,
      username: normalUsername,
    });
  };

  return {
    state,
    close,
    onChange,
    onSignUp,
  };
}
