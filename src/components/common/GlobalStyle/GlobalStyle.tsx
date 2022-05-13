import { css, Global } from '@emotion/react';
import { palette, themePalette, themes } from '../../../styles';

interface GlobalStyleProps {}

function GlobalStyle(props: GlobalStyleProps) {
  return <Global styles={globalStyle} />;
}

const globalStyle = css`
  html,
  body,
  #root {
    margin: 0;
    padding: 0;
    height: 100%;
    color: ${themePalette.text1};
    background: ${themePalette.bg4};
  }

  header,
  footer {
    background: ${themePalette.bg4};
  }

  html {
    box-sizing: border-box;
    * {
      box-sizing: inherit;
      font-family: 'Montserrat', sans-serif, -apple-system, BlinkMacSystemFont,
        'Helvetica Neue', 'Apple SD Gothic Neo', 'Malgun Gothic', '맑은 고딕',
        나눔고딕, 'Nanum Gothic', 'Noto Sans KR', 'Noto Sans CJK KR', arial,
        돋움, Dotum, Tahoma, Geneva, sans-serif;
      text-rendering: optimizeLegibility;
      -webkit-font-smoothing: subpixel-antialiased;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
    }
  }

  body {
    ${themes.light}
  }

  @media (prefers-color-scheme: dark) {
    body {
      ${themes.dark}
    }
  }

  body[data-theme='light'] {
    ${themes.light};
  }

  body[data-theme='dark'] {
    ${themes.dark};
  }

  h1,
  h2,
  h3,
  h4,
  h5,
  h6,
  p,
  span,
  a {
    color: ${themePalette.text1};
    ::selection {
      color: ${palette.white};
      background: ${palette.blue['600']};
    }
  }
  a {
    text-decoration: none;
    cursor: pointer;
  }
`;

export default GlobalStyle;
