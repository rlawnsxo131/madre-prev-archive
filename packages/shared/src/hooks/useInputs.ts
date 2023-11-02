import type { ChangeEvent } from 'react';
import { useCallback, useState } from 'react';

export function useInputs(initialState: Record<string, string>) {
  const [state, setState] = useState(initialState);

  const onChange = useCallback((e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setState((prev) => ({
      ...prev,
      [name]: value,
    }));
  }, []);

  return { state, onChange };
}
