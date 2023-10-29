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
    flexBasis: '',
    flexDirection: 'row',
    flexFlow: '',
    flexGrow: '',
    flexShrink: '',
    justifyContent: '',
    alignContent: '',
    alignItems: '',
    alignSelf: '',
    className: '',
  },
  render: (args) => {
    const itemStyle = {
      padding: '1rem',
      border: '1px solid black',
      display: 'flex',
      flex: '0',
    };

    return (
      <FlexLayout {...args}>
        <div style={itemStyle}>item1</div>
        <div style={itemStyle}>item2</div>
      </FlexLayout>
    );
  },
};
