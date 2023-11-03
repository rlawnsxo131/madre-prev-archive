import { AnimatePresence, motion } from 'framer-motion';
import type { PropsWithoutRef, ReactNode } from 'react';
import { forwardRef, useState } from 'react';
import { createPortal } from 'react-dom';

import { useIsomorphicLayoutEffect } from '@/hooks';
import { getElement } from '@/lib/utils/getElement';

import { Overlay } from '../../Overlay';
import { block, content } from './ModalLayout.css';

export type ModalLayoutProps = PropsWithoutRef<{
  children?: ReactNode;
  visible: boolean;
  duration?: number;
  container?: Parameters<typeof getElement>[0];
}>;

export const ModalLayout = forwardRef<HTMLDivElement, ModalLayoutProps>(
  function ({ children, visible, duration = 0.15, container = null }, ref) {
    const [mountNode, setMountNode] =
      useState<ReturnType<typeof getElement>>(null);

    useIsomorphicLayoutEffect(() => {
      setMountNode(getElement(container) || document.body);
    }, [container]);

    return mountNode
      ? createPortal(
          <>
            <Overlay visible={visible} duration={duration} />
            <AnimatePresence initial={false}>
              {visible && (
                <div className={block} ref={ref}>
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
  },
);
