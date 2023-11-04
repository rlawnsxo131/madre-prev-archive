import type { ChangeEvent } from 'react';
import { useCallback, useState } from 'react';

export function useInputs<
  Value extends Record<string, string> = Record<string, string>,
>(defaultValue: Value) {
  const [state, setState] = useState(defaultValue);

  const onChange = useCallback((e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setState((prev) => ({
      ...prev,
      [name]: value,
    }));
  }, []);

  return {
    state,
    onChange,
  };
}
