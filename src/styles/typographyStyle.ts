import { css } from '@emotion/react';
import media from './media';

export default css`
  h1,
  h2,
  h3,
  h4,
  h5,
  h6,
  p,
  span,
  a {
    white-space: pre-line;
    word-break: break-word;
    overflow-wrap: break-word;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    line-height: 1.5;
  }

  ${media.xxxsmall} {
    h1 {
      font-size: 1.5rem;
    }
    h2 {
      font-size: 1.25rem;
    }
    h3 {
      font-size: 1.125rem;
    }
    h4 {
      margin: 1rem 0;
      font-size: 1rem;
    }
    h5 {
      margin: 0.875rem 0;
      font-size: 0.875rem;
    }
    h6 {
      margin: 0.85rem 0;
      font-size: 0.85rem;
    }
    p {
      font-size: 1rem;
    }
    li {
      font-size: 0.9rem;
    }
  }
  ${media.small} {
    h1 {
      font-size: 1.85rem;
    }
    h2 {
      font-size: 1.5rem;
    }
    h3 {
      font-size: 1.25rem;
    }
    h4 {
      margin: 1.125rem 0;
      font-size: 1.125rem;
    }
    h5 {
      margin: 1rem 0;
      font-size: 1rem;
    }
    h6 {
      margin: 0.875rem 0;
      font-size: 0.875rem;
    }
    p {
      font-size: 1rem;
    }
    li {
      font-size: 1rem;
    }
  }
  ${media.medium} {
    h1 {
      font-size: 2rem;
    }
    h2 {
      font-size: 1.825rem;
    }
    h3 {
      font-size: 1.5rem;
    }
    h4 {
      margin: 1.25rem 0;
      font-size: 1.25rem;
    }
    h5 {
      margin: 1.125rem 0;
      font-size: 1.125rem;
    }
    h6 {
      margin: 1rem 0;
      font-size: 1rem;
    }
    p {
      font-size: 1.125rem;
    }
  }
`;
