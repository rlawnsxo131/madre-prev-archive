import { css } from '@emotion/react';
import { palette, transitions } from '../../../styles';
import { themePalette } from '../../../styles/themePalette';

const shakeAnimation = css`
  animation: ${transitions.shake} 0.5s 0.25s ease-in-out;
`;

const rightBlock = css`
  flex: 2 1 0;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 1rem 1rem 1.5rem 1rem;
  border-radius: 0.25rem;
`;

const rightBlockHeader = css`
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1rem;

  h1 {
    margin: 0;
    &.is-error {
      color: ${palette.pink['600']};
      ${shakeAnimation}
    }
  }

  .popup-auth-close-icon {
    align-self: flex-end;
    width: 1.25rem;
    height: 1.25rem;
    color: ${themePalette.fill1};
    cursor: pointer;
  }
`;

const rightBlockBody = css`
  display: flex;
  flex-direction: column;
  svg {
    width: 100%;
    height: auto;
  }
`;

export default {
  shakeAnimation,
  rightBlock,
  rightBlockBody,
  rightBlockHeader,
};
