import { renderHook } from '@testing-library/react-hooks';
import useTransitionTimeoutEffect from '../useTransitionTimeoutEffect';

describe('useTransitionTimeoutEffect', () => {
  beforeAll(() => {
    jest.useFakeTimers();
  });

  afterEach(() => {
    jest.clearAllTimers();
  });

  it('closed is false', async () => {
    const { result } = renderHook(() =>
      useTransitionTimeoutEffect({ visible: true }),
    );

    jest.advanceTimersByTime(250);

    expect(result.current).toBe(false);
  });

  it('closed is true', async () => {
    const { result } = renderHook(() =>
      useTransitionTimeoutEffect({ visible: false }),
    );

    jest.advanceTimersByTime(250);

    expect(result.current).toBe(true);
  });
});
