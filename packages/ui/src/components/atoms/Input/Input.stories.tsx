import { useInputs } from '@madre/react';
import type { Meta, StoryObj } from '@storybook/react';

import { FlexLayout } from '../FlexLayout';
import type { InputProps } from './Input';
import { Input } from './Input';

const meta = {
  title: 'atoms/Input',
  component: Input,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof Input>;

export default meta;
type Story = StoryObj<typeof meta>;

function DefaultComponent(args: Story['args']) {
  const {
    state: { state1, state2, state3 },
    onChange,
  } = useInputs({ state1: '', state2: '', state3: '' });

  const itemBoxStyle = {
    marginTop: '1rem',
  };

  const _args: InputProps = {
    style: { marginTop: '1rem' },
  };

  return (
    <FlexLayout flexDirection="column">
      <FlexLayout justifyContent="center">
        <Input {...args} name="name" id="name" />
      </FlexLayout>

      {/* default */}
      <FlexLayout flexDirection="column" style={itemBoxStyle}>
        <Input
          size="small"
          name="state1"
          value={state1}
          onChange={onChange}
          {..._args}
        />
        <Input
          size="medium"
          name="state2"
          value={state2}
          onChange={onChange}
          {..._args}
        />
        <Input
          size="large"
          name="state3"
          value={state3}
          onChange={onChange}
          {..._args}
        />
      </FlexLayout>

      {/* warn */}
      <FlexLayout flexDirection="column" style={itemBoxStyle}>
        <Input
          size="small"
          name="state1"
          value={state1}
          onChange={onChange}
          error
          {..._args}
        />
        <Input
          size="medium"
          name="state2"
          value={state2}
          onChange={onChange}
          error
          {..._args}
        />
        <Input
          size="large"
          name="state3"
          value={state3}
          onChange={onChange}
          error
          {..._args}
        />
      </FlexLayout>
    </FlexLayout>
  );
}

export const Default: Story = {
  args: {
    error: false,
    size: 'medium',
    value: 'value',
  },
  render: (args) => <DefaultComponent {...args} />,
};
