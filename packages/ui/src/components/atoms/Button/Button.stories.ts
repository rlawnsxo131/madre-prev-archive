import type { Meta, StoryObj } from '@storybook/react';

import { Button } from './Button';

const meta = {
  title: 'atoms/Button',
  component: Button,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof Button>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
  args: {
    children: 'button',
    theme: 'primary',
    size: 'medium',
    radius: 'medium',
    onClick: () => alert('click'),
  },
  // play: async ({ canvasElement }) => {
  //   const canvas = within(canvasElement);
  //   await userEvent.click(canvas.getByRole('button'));
  // },
};

export const Secondary: Story = {
  args: {
    children: 'button',
    theme: 'secondary',
    size: 'medium',
    radius: 'medium',
  },
};

export const Warn: Story = {
  args: {
    children: 'button',
    size: 'medium',
    theme: 'warn',
    radius: 'medium',
  },
};
