import { css } from '@emotion/react';
import { media, palette } from '../../styles';

export const navigationActiveColor = palette.pink['700'];

export const navigationDisplay = css`
  ${media.xxxsmall} {
    display: none;
  }
  ${media.small} {
    display: flex;
  }
`;
export const navigationMobileDisplay = css`
  ${media.small} {
    display: none;
  }
`;
