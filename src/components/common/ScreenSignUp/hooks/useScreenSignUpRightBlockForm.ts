import React from 'react';
import { useEffect, useRef, useState } from 'react';
import useScreenSignUpActions from '../../../../hooks/screenSignUp/useScreenSignUpActions';
import useScreenSignUpState from '../../../../hooks/screenSignUp/useScreenSignUpState';
import useToast from '../../../../hooks/useToast';
import { isNormalEnglishString, normalizeString } from '../../../../lib/utils';
import authApi from '../../../../store/api/authApi';

export default function useScreenSignUpRightBlockForm() {
  const { isError, isValidateError, isConflictError, access_token } =
    useScreenSignUpState();
  const [googleSignUp] = authApi.usePostGoogleSignUpMutation();
  const { error, warn } = useToast();
  const {
    close,
    setIsValidateError,
    resetIsValidateError,
    resetIsConflictError,
  } = useScreenSignUpActions();

  const inputRef = useRef<HTMLInputElement>(null);
  const [username, setUsername] = useState('');

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    resetIsValidateError();
    resetIsConflictError();
    setUsername(e.target.value);
  };

  const onSignUp = async (e: React.FormEvent) => {
    e.preventDefault();
    const normalizedAccessToken = normalizeString(access_token);
    const normalizedUsername = normalizeString(username);

    if (!normalizedAccessToken || !normalizedUsername) {
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
    if (!isNormalEnglishString(normalizedUsername)) {
      setIsValidateError();
      return;
    }
    await googleSignUp({
      access_token: normalizedAccessToken,
      username: normalizedUsername,
    });
  };

  useEffect(() => {
    if (!isValidateError) return;
    inputRef.current?.focus();
    warn('이름을 다시 확인해 주세요.(영문, 숫자 1~16자)', 'top-center');
  }, [isValidateError, warn]);

  useEffect(() => {
    if (!isConflictError) return;
    inputRef.current?.focus();
    error('중복된 이름입니다.', 'top-center');
  }, [isConflictError, error]);

  useEffect(() => {
    if (!isError) return;
    error('에러가 발생했습니다. 잠시후 다시 시도해 주세요.', 'top-center');
  }, [isError, error]);

  return {
    inputRef,
    username,
    onChange,
    close,
    onSignUp,
  };
}
