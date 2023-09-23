import { render } from '@testing-library/react';
import { renderHook } from '@testing-library/react-hooks';
import { MemoryRouter } from 'react-router-dom';
import App from '../../App';
import useIsUserPath from '../useIsUserPath';

describe('useIsUserPath', () => {
  it('isUserPath', () => {
    render(
      <MemoryRouter initialEntries={['/@userpath']}>
        <App />
      </MemoryRouter>,
    );
    const { result } = renderHook(() => useIsUserPath());

    console.log(result.current);

    expect(result.current).toBe(true);
  });
});
