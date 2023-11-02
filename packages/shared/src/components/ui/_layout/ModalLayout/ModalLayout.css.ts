import { style } from '@vanilla-extract/css';

import { themes, zIndices } from '@/styles';

export const block = style({
  position: 'fixed',
  top: 0,
  left: 0,
  width: '100%',
  height: '100%',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  zIndex: zIndices.modal,
  backgroundColor: 'transparent',
});

export const content = style({
  position: 'relative',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  borderRadius: '0.25rem',
  backgroundColor: themes.color['bg-element1'],
  boxShadow: '0 2px 12px 0 rgba(0, 0, 0, 0.09)',
});
