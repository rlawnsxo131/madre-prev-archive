import { renderHook, act } from '@testing-library/react-hooks';
import useInputs from '../useInputs';

describe('useInputs', () => {
  const initialValue = {
    name: 'name',
  };

  it('onChange', () => {
    const { result } = renderHook(() => useInputs(initialValue));
    const event = {
      target: { name: 'name', value: 'changename' },
    } as React.ChangeEvent<HTMLInputElement>;

    expect(result.current.state.name).toBe(initialValue.name);

    // 현재 테스트는 잘 통과하는데 18 을 제대로 지원 안하는듯..?
    act(() => result.current.onChange(event));

    expect(result.current.state.name).toBe(event.target.value);
  });

  it('onReset', () => {
    const { result } = renderHook(() => useInputs(initialValue));

    expect(result.current.state.name).toBe(initialValue.name);

    act(() => result.current.onReset());

    expect(result.current.state.name).toBe('');
  });
});
