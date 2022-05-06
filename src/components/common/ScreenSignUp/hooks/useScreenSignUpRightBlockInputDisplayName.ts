import { useState } from 'react';
import useScreenSignUpActions from '../../../../hooks/screenSignUp/useScreenSignUpActions';
import useScreenSignUpState from '../../../../hooks/screenSignUp/useScreenSignUpState';
import useInputs from '../../../../hooks/useInputs';
import { isNormalEnglishString, normalizeString } from '../../../../lib/utils';
import authApi from '../../../../store/api/authApi';

export default function useScreenSignUpInputDisplayName() {
  const { isError, access_token } = useScreenSignUpState();
  const [googleSignUp] = authApi.usePostGoogleSignUpMutation();
  const { close } = useScreenSignUpActions();
  const { state, onChange } = useInputs({
    display_name: '',
  });
  const [isValidateError, setIsValidateError] = useState(true);

  const onSignUp = async () => {
    const normalizedAccessToken = normalizeString(access_token);
    const normalizedDisplayName = normalizeString(state.display_name);

    if (!normalizedAccessToken || !normalizedDisplayName) {
      if (!normalizedAccessToken) {
        console.log('system error ');
        return;
      }
      setIsValidateError(true);
      return;
    }
    if (!isNormalEnglishString(normalizedDisplayName)) {
      setIsValidateError(true);
      return;
    }
    await googleSignUp({
      access_token: normalizedAccessToken,
      display_name: normalizedDisplayName,
    });
  };

  return {
    state,
    isError,
    isValidateError,
    close,
    onChange,
    onSignUp,
  };
}
