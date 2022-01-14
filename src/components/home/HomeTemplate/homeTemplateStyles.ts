import { css } from '@emotion/react';

const itemBlock = css`
  display: flex;
  justify-content: center;
  align-items: center;
  /* &:not(:nth-of-type(3)) {
    padding: 0 0.5rem;
  } */
  &:nth-of-type(1) {
    padding: 0 0.5rem;
  }
  &:nth-of-type(2) {
    padding-left: 0.5rem;
  }
`;

export default {
  itemBlock,
};
