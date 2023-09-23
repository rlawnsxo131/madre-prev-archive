import { useReducer, useCallback } from 'react';

interface UseInputsAction {
  name: string;
  value: string;
}

function reducer<T>(state: T, action: UseInputsAction | null) {
  if (!action) {
    const initialState: Record<string, any> = {};
    Object.keys(state).forEach((key) => {
      initialState[key] = '';
    });
    return initialState;
  }
  return {
    ...state,
    [action.name]: action.value,
  };
}

export default function useInputs<T>(defaultValues: T) {
  const [state, dispatch] = useReducer(reducer, defaultValues);

  const onChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
      dispatch({
        name: e.target.name,
        value: e.target.value,
      });
    },
    [dispatch],
  );

  const onReset = useCallback(() => {
    dispatch(null);
  }, [dispatch]);

  return {
    state,
    onChange,
    onReset,
    dispatch,
  };
}
