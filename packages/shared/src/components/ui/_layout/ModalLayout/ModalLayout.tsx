import { AnimatePresence, motion } from 'framer-motion';
import type { PropsWithChildren } from 'react';
import { useState } from 'react';
import { createPortal } from 'react-dom';

import { useIsomorphicLayoutEffect } from '@/hooks';
import { getContainer } from '@/lib/utils/getContainer';

import { Overlay } from '../../Overlay';
import { block, content } from './ModalLayout.css';

export type ModalLayoutProps = PropsWithChildren<{
  visible: boolean;
  duration?: string;
  container?: Parameters<typeof getContainer>[0];
}>;

export function ModalLayout({
  children,
  visible,
  duration = '.15',
  container = null,
}: ModalLayoutProps) {
  const [mountNode, setMountNode] =
    useState<ReturnType<typeof getContainer>>(null);

  useIsomorphicLayoutEffect(() => {
    setMountNode(getContainer(container) || document.body);
  }, [container]);

  return mountNode
    ? createPortal(
        <>
          <Overlay visible={visible} />
          <AnimatePresence initial={false}>
            {visible && (
              <div className={block}>
                <motion.div
                  className={content}
                  initial={{
                    opacity: 0,
                    scale: 0.75,
                  }}
                  animate={{
                    opacity: 1,
                    scale: 1,
                  }}
                  exit={{
                    opacity: 0,
                    scale: 0.75,
                  }}
                  transition={{
                    duration,
                    ease: 'easeIn',
                  }}
                >
                  {children}
                </motion.div>
              </div>
            )}
          </AnimatePresence>
        </>,
        mountNode,
      )
    : mountNode;
}
