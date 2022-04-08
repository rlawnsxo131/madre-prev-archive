import useScreenSignUpActions from '../../../../hooks/screenSignUp/useScreenSignUpActions';
import useScreenSignUpState from '../../../../hooks/screenSignUp/useScreenSignUpState';
import useInputs from '../../../../hooks/useInputs';
import { normalizedString } from '../../../../lib/utils';
import authApi from '../../../../store/api/authApi';

export default function useScreenSignUpRightBlock() {
  const { access_token } = useScreenSignUpState();
  const [googleSignUp] = authApi.usePostGoogleSignUpMutation();
  const { close } = useScreenSignUpActions();
  const { state, onChange } = useInputs({
    display_name: '',
  });

  const onSignUp = async () => {
    const normalAccessToken = normalizedString(access_token);
    const normalDisplayName = normalizedString(state.display_name);

    if (!normalAccessToken || !normalDisplayName) {
      if (!normalAccessToken) {
        console.log('시스템 에러');
        return;
      }
      return;
    }
    await googleSignUp({
      access_token: normalAccessToken,
      display_name: normalDisplayName,
    });
  };

  return {
    state,
    close,
    onChange,
    onSignUp,
  };
}
