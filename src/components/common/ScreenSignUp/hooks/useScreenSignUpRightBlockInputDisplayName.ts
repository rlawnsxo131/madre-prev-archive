import { useEffect, useRef, useState } from 'react';
import useScreenSignUpActions from '../../../../hooks/screenSignUp/useScreenSignUpActions';
import useScreenSignUpState from '../../../../hooks/screenSignUp/useScreenSignUpState';
import useToast from '../../../../hooks/useToast';
import { isNormalEnglishString, normalizeString } from '../../../../lib/utils';
import authApi from '../../../../store/api/authApi';

export default function useScreenSignUpInputDisplayName() {
  const { isError, isValidateError, access_token } = useScreenSignUpState();
  const [googleSignUp] = authApi.usePostGoogleSignUpMutation();
  const { error, warn } = useToast();
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
        error(
          '에러가 발생했습니다. 화면을 닫고 로그인을 다시 시도해 주세요.',
          'top-center',
        );
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
    warn('이름을 다시 확인해 주세요.(영문, 숫자 1~16자)', 'top-center');
  }, [isValidateError]);

  useEffect(() => {
    if (!isError) return;
    error('에러가 발생했습니다. 잠시후 다시 시도해 주세요.', 'top-center');
  }, [isError]);

  return {
    inputRef,
    displayName,
    onChange,
    close,
    onSignUp,
  };
}
