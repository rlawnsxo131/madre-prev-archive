import useScreenSignUpActions from '../../../../hooks/screenSignUp/useScreenSignUpActions';
import useScreenSignUpState from '../../../../hooks/screenSignUp/useScreenSignUpState';
import useInputs from '../../../../hooks/useInputs';
import { normalizeString } from '../../../../lib/utils';
import authApi from '../../../../store/api/authApi';

export default function useScreenSignUpRightBlockInputNickname() {
  const { access_token } = useScreenSignUpState();
  const [googleSignUp] = authApi.usePostGoogleSignUpMutation();
  const { close } = useScreenSignUpActions();
  const { state, onChange } = useInputs({
    display_name: '',
  });

  const onSignUp = async () => {
    const normalizedAccessToken = normalizeString(access_token);
    const normalizedDisplayName = normalizeString(state.display_name);

    if (!normalizedAccessToken || !normalizedDisplayName) {
      if (!normalizedAccessToken) {
        console.log('시스템 에러');
        return;
      }
      console.log('displayName 에러');
      return;
    }
    await googleSignUp({
      access_token: normalizedAccessToken,
      display_name: normalizedDisplayName,
    });
  };

  return {
    state,
    close,
    onChange,
    onSignUp,
  };
}
