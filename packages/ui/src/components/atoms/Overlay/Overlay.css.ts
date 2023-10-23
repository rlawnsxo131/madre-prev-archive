import { style } from '@vanilla-extract/css';

import { vars } from '@/styles';

export const block = style({
  position: 'fixed',
  left: 0,
  top: 0,
  width: '100%',
  height: '100%',
  background: vars.color.overlay,
  zIndex: vars.zIndices.overlay,
  selectors: {
    '&': {
      '@supports': {
        '(-webkit-touch-callout: none)': {
          height: '-webkit-fill-available',
        },
      },
    },
  },
});
