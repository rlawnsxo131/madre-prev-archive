import { useCallback, useState } from 'react';

export function useBooleanState(defaultValue: boolean | (() => boolean)) {
  const [value, setValue] = useState(
    typeof defaultValue === 'function' ? defaultValue() : defaultValue,
  );

  const onSetValueToTrue = useCallback(() => {
    setValue(true);
  }, []);

  const onSetValueToFalse = useCallback(() => {
    setValue(false);
  }, []);

  return {
    value,
    onSetValueToTrue,
    onSetValueToFalse,
  };
}
