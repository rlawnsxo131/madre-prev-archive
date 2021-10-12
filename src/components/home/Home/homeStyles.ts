import { css } from '@emotion/react';
import media from '../../../styles/media';

export const homeBlock = css`
  display: flex;
  gap: 1.5rem;
  & + & {
    margin-top: 2rem;
  }
  ${media.xxxsmall} {
    flex-direction: column;
  }
  ${media.medium} {
    flex-flow: row;
    justify-content: space-around;
  }
`;

export const homeBlockItemCommon = css`
  padding: 4rem 0 3rem 0;
`;

export const homeH3 = css`
  margin: 0;
  padding: 0;
  font-size: 1.75rem;
  font-weight: 550;
  line-height: 2;
  white-space: pre-line;
`;

export const homeH5 = css`
  margin: 0;
  padding: 0;
  font-size: 1rem;
`;

export const homeP = css`
  margin: 0;
  padding: 0;
  font-size: 1.125rem;
  font-weight: 550;
  line-height: 2;
  white-space: pre-line;
`;
