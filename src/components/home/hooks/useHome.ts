import { useCallback, useState } from 'react';
import { useRecoilValue } from 'recoil';
import { darkmodeStateSelector } from '../../../atoms/darkmodeState';

export default function useHome() {
  const { theme } = useRecoilValue(darkmodeStateSelector);
  const [visible, setVisible] = useState<boolean>(false);

  const handleVisible = useCallback(() => {
    setVisible((state) => !state);
  }, []);

  return {
    theme,
    visible,
    handleVisible,
  };
}
