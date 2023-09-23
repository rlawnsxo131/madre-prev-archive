import { render } from '@testing-library/react';
import Button from '../common/Button';

describe('<Button />', () => {
  it('render test', () => {
    const text = 'button';
    const result = render(<Button>{text}</Button>);

    const renderText = result.getByText(text);
    expect(renderText).toBeInTheDocument();
  });
});
