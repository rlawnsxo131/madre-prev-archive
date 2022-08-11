import 'jest';
import { renderHook, act } from '@testing-library/react-hooks';
import useInputs from './useInputs';

describe('hooks', () => {
  test('useInputs', async () => {
    const { result } = renderHook(() => useInputs({ name: 'name' }));
    const event = {
      preventDefault() {},
      target: { name: 'name', value: 'changename' },
    } as React.ChangeEvent<HTMLInputElement>;

    expect(result.current.state.name).toBe('name');

    act(() => result.current.onChange(event));

    expect(result.current.state.name).toBe('changename');
  });
});
