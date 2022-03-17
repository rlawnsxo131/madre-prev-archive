import { css } from '@emotion/react';
import { media, palette } from '../../../styles';
import { themePalette } from '../../../styles/themePalette';

const section = css`
  display: flex;
  margin-top: 1.5rem;

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
    color: ${themePalette.text2};
    font-weight: 600;
    line-height: 2;
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
    gap: 2rem;
  }
`;

const itemCommon = css`
  ${media.xxxsmall} {
    padding: 3rem 0 0 0;
  }
  ${media.medium} {
    padding: 3rem 0;
  }
`;

const descriptionBlock = css`
  ${itemCommon};
  flex: 1;
  display: flex;
`;

const description = css`
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  gap: 1.25rem;
  h2,
  h3,
  p {
    letter-spacing: 0.1rem;
    ::selection {
      color: ${palette.white};
      background: ${palette.pink['600']};
    }
  }
`;

const imageBlock = css`
  ${itemCommon};
  flex: 1;
  display: flex;
  justify-content: center;
  gap: 1.25rem;
  svg {
    max-width: 35rem;
    height: auto;
  }
`;

export default {
  section,
  itemCommon,
  descriptionBlock,
  description,
  imageBlock,
};
