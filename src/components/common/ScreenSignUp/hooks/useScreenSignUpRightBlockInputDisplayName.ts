import { useEffect, useRef, useState } from 'react';
import useScreenSignUpActions from '../../../../hooks/screenSignUp/useScreenSignUpActions';
import useScreenSignUpState from '../../../../hooks/screenSignUp/useScreenSignUpState';
import { isNormalEnglishString, normalizeString } from '../../../../lib/utils';
import authApi from '../../../../store/api/authApi';

export default function useScreenSignUpInputDisplayName() {
  const { isError, isValidateError, access_token } = useScreenSignUpState();
  const [googleSignUp] = authApi.usePostGoogleSignUpMutation();
  const { close, setIsValidateError, resetIsValidateError } =
    useScreenSignUpActions();

  const inputRef = useRef<HTMLInputElement>(null);
  const [displayName, setDisplayName] = useState('');

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    resetIsValidateError();
    setDisplayName(e.target.value);
  };

  const onSignUp = async () => {
    const normalizedAccessToken = normalizeString(access_token);
    const normalizedDisplayName = normalizeString(displayName);

    if (!normalizedAccessToken || !normalizedDisplayName) {
      if (!normalizedAccessToken) {
        console.log('system error');
        return;
      }
      setIsValidateError();
      return;
    }
    if (!isNormalEnglishString(normalizedDisplayName)) {
      setIsValidateError();
      return;
    }
    await googleSignUp({
      access_token: normalizedAccessToken,
      display_name: normalizedDisplayName,
    });
  };

  useEffect(() => {
    if (!isValidateError) return;
    inputRef.current?.focus();
  }, [isValidateError]);

  return {
    inputRef,
    displayName,
    onChange,
    close,
    onSignUp,
  };
}
