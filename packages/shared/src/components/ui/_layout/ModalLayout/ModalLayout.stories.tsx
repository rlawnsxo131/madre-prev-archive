import type { Meta, StoryObj } from '@storybook/react';

import { FlexLayout } from '../FlexLayout';
import { ModalLayout } from './ModalLayout';

const meta = {
  title: '_layout/ModalLayout',
  component: ModalLayout,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof ModalLayout>;

export default meta;
type Story = StoryObj<typeof meta>;

function DefaultComponent({ visible }: Pick<Story['args'], 'visible'>) {
  return (
    <>
      <ModalLayout visible={visible}>
        <FlexLayout flexDirection="column" style={{ padding: '0.5rem 1rem' }}>
          <h1>ModalLayout</h1>
        </FlexLayout>
      </ModalLayout>
    </>
  );
}

export const Default: Story = {
  args: {
    visible: false,
  },
  render: (args) => <DefaultComponent {...args} />,
};
