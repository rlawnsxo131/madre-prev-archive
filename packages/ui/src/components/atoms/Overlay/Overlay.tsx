import { AnimatePresence, motion } from 'framer-motion';
import { type ComponentProps } from 'react';

import { block } from './Overlay.css';

export type OverlayProps = ComponentProps<typeof motion.div> & {
  visible: boolean;
  duration?: string;
};

export function Overlay({ visible, duration = '.15', ...props }: OverlayProps) {
  return (
    <AnimatePresence initial={false}>
      {visible && (
        <motion.div
          className={block}
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          exit={{ opacity: 0 }}
          transition={{
            duration,
            ease: 'easeIn',
          }}
          {...props}
        />
      )}
    </AnimatePresence>
  );
}
