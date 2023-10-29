import type { Meta, StoryObj } from '@storybook/react';
import { useState } from 'react';

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
  const [state, setState] = useState('');

  return (
    <Input value={state} onChange={(e) => setState(e.target.value)} {...args} />
  );
}

export const Default: Story = {
  args: {},
  render: (args) => <DefaultComponent {...args} />,
};
