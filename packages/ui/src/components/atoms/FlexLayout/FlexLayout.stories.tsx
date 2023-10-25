import type { Meta, StoryObj } from '@storybook/react';

import { FlexLayout } from './FlexLayout';

const meta = {
  title: 'atoms/FlexLayout',
  component: FlexLayout,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof FlexLayout>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    flex: 1,
    flexBasis: '',
    flexDirection: 'column',
    flexFlow: '',
    flexGrow: '',
    flexShrink: '',
    justifyContent: '',
    alignContent: '',
    alignItems: '',
    alignSelf: '',
    className: '',
  },
  render: (args) => (
    <FlexLayout {...args}>
      <div>flex</div>
      <div>box</div>
    </FlexLayout>
  ),
};
