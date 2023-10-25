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
    variant: 'solid',
    full: false,
  },
};

export const Secondary: Story = {
  args: {
    children: 'button',
    theme: 'secondary',
    size: 'medium',
    radius: 'medium',
    variant: 'solid',
    full: false,
  },
};

export const Warn: Story = {
  args: {
    children: 'button',
    theme: 'warn',
    size: 'medium',
    radius: 'medium',
    variant: 'solid',
    full: false,
  },
};
