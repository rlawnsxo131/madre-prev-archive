import classNames from 'classnames';
import { AnimatePresence, motion } from 'framer-motion';
import type { ComponentProps, PropsWithoutRef } from 'react';
import { forwardRef } from 'react';

import { block } from './Overlay.css';

/**
 * @description ...props 의 형태로 motion.div 의 모든 props 를 받는 경우
 * 아래 경고가 발생하므로 필요할때 props 요소를 하나씩 추가 정의
 *
 * react-dom.development.js:86 Warning: Unknown event handler property `onPointerEnterCapture`. It will be ignored.
 */
export type OverlayProps = PropsWithoutRef<
  ComponentProps<typeof motion.div> & {
    visible: boolean;
    duration?: string;
  }
>;

export const Overlay = forwardRef<HTMLDivElement, OverlayProps>(function (
  { className, visible, duration = '.15' },
  ref,
) {
  return (
    <AnimatePresence initial={false}>
      {visible && (
        <motion.div
          ref={ref}
          className={classNames(block, className)}
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          exit={{ opacity: 0 }}
          transition={{
            duration,
            ease: 'easeIn',
          }}
        />
      )}
    </AnimatePresence>
  );
});
