import { style } from '@vanilla-extract/css';

import { themes, zIndices } from '@/styles';

export const block = style({
  position: 'fixed',
  left: 0,
  top: 0,
  width: '100%',
  height: '100%',
  backgroundColor: themes.color['overlay'],
  zIndex: zIndices['overlay'],
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
