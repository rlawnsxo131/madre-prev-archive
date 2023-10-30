import type { Meta, StoryObj } from '@storybook/react';

import { Popup } from './Popup';

const meta = {
  title: 'features/Popup',
  component: Popup,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof Popup>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    visible: false,
    duration: '.15',
  },
  render: (args) => (
    <div style={{ height: '320px' }}>
      <Popup {...args} />
    </div>
  ),
};
