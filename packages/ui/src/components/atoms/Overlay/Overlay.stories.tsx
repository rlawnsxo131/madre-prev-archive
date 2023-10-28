import type { Meta, StoryObj } from '@storybook/react';

import { Overlay } from './Overlay';

const meta = {
  title: 'atoms/Overlay',
  component: Overlay,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof Overlay>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    visible: false,
    duration: '.15',
  },
  render: (args) => (
    <div style={{ width: '320px', height: '320px' }}>
      <Overlay {...args} />
    </div>
  ),
};
