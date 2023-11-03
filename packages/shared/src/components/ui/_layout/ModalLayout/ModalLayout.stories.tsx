import type { Meta, StoryObj } from '@storybook/react';
import type { BaseSyntheticEvent } from 'react';
import { useState } from 'react';

import { useRefEffect } from '@/hooks';

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

  const ref = useRefEffect<HTMLDivElement>((div) => {
    const handler = (e: BaseSyntheticEvent | Event | TouchEvent) => {
      if (!div.contains(e.target)) return;
      setVisible(false);
    };

    document.addEventListener('mousedown', handler);
    document.addEventListener('touchstart', handler);

    return () => {
      document.removeEventListener('mousedown', handler);
      document.removeEventListener('touchstart', handler);
    };
  }, []);

  return (
    <>
      <Button onClick={() => setVisible(true)}>click</Button>
      <ModalLayout visible={visible} ref={ref}>
        <FlexLayout flexDirection="column" style={{ padding: '0.5rem 1rem' }}>
          <h1>ModalLayout</h1>
          <button onClick={() => setVisible(false)}>ok</button>
        </FlexLayout>
      </ModalLayout>
    </>
  );
}

export const Default: Omit<Story, 'args'> = {
  render: () => <DefaultComponent />,
};
