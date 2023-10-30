import { Portal } from '@madre/react';
import type { ComponentProps } from 'react';

import { Overlay } from '@/components/atoms/Overlay';

export type PopupProps = ComponentProps<typeof Overlay>;

export function Popup({ visible, duration }: PopupProps) {
  return (
    <Portal>
      <Overlay visible={visible} duration={duration} />
      <div>content</div>
    </Portal>
  );
}
