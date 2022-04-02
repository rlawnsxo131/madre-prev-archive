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
    const normalaccess_token = normalizedString(access_token);
    const normalDisplayName = normalizedString(state.display_name);

    if (!normalaccess_token || !normalDisplayName) {
      if (!normalaccess_token) {
        console.log('시스템 에러');
        return;
      }
      return;
    }
    await googleSignUp({
      access_token: normalaccess_token,
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
