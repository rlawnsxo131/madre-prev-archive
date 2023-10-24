import type { Meta, StoryObj } from '@storybook/react';

import { FlexBox } from './FlexBox';

const meta = {
  title: 'atoms/FlexBox',
  component: FlexBox,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof FlexBox>;

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
    <FlexBox {...args}>
      <div>flex</div>
      <div>box</div>
    </FlexBox>
  ),
};
