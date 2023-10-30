import type { PropsWithChildren, ReactNode } from 'react';
import { useState } from 'react';
import { createPortal } from 'react-dom';

import { useIsomorphicLayoutEffect } from '@/hooks';

function getContainer(container: PortalProps['container']) {
  return typeof container === 'function' ? container() : container;
}

export type PortalProps = PropsWithChildren<{
  children?: ReactNode;
  container?: Element | (() => Element | null) | null;
}>;

export function Portal({ children, container }: PortalProps) {
  const [mountNode, setMountNode] =
    useState<ReturnType<typeof getContainer>>(null);

  useIsomorphicLayoutEffect(() => {
    setMountNode(getContainer(container) || document.body);
  }, [container]);

  return mountNode ? createPortal(children, mountNode) : mountNode;
}
