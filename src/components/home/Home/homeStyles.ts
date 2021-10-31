import { css } from '@emotion/react';
import media from '../../../styles/media';
import palette from '../../../styles/palette';

export const homeBlock = css`
  display: flex;
  h1,
  h2,
  h3,
  h4,
  p {
    margin: 0;
    padding: 0;
    white-space: pre-line;
    line-height: 1.5;
  }
  p {
    font-size: 1.125rem;
    color: ${palette.gray['700']};
    font-weight: 600;
    line-height: 2;
  }
  & + & {
    margin-top: 2rem;
  }
  ${media.xxxsmall} {
    gap: 0;
    flex-direction: column;
  }
  ${media.medium} {
    &:nth-of-type(odd) {
      flex-direction: row;
    }
    &:nth-of-type(even) {
      flex-direction: row-reverse;
    }
    justify-content: space-around;
    gap: 1.5rem;
  }
`;

export const homeBlockItemCommon = css`
  padding: 4rem 0 3rem 0;
`;
