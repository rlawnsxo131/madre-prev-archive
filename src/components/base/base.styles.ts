import { css } from '@emotion/react';
import { media, mediaQuery } from '../../styles';

const responsive = css`
  ${media.xxxsmall} {
    width: 93%;
  }
  ${media.medium} {
    width: 96%;
  }
  ${mediaQuery(1285)} {
    max-width: 1250px;
  }
`;

export default {
  responsive,
};
