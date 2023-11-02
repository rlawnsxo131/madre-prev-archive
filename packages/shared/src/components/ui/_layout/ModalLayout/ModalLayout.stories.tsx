import type { Meta, StoryObj } from '@storybook/react';
import { useState } from 'react';

import { Button } from '../../Button';
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

function DefaultComponent() {
  const [visible, setVisible] = useState(false);

  return (
    <>
      <Button onClick={() => setVisible(true)}>click</Button>
      <ModalLayout visible={visible}>
        <FlexLayout flexDirection="column" style={{ padding: '0.5rem 1rem' }}>
          <div>content</div>
          <div>
            <button onClick={() => setVisible(false)}>ok</button>
          </div>
        </FlexLayout>
      </ModalLayout>
    </>
  );
}

export const Default: Omit<Story, 'args'> = {
  render: () => <DefaultComponent />,
};
