import { css, Global } from '@emotion/react';
import { ColorTheme, useColorThemeValue } from '../../../atoms/colorThemeState';
import { themeColor } from '../../../styles';

interface GlobalStyleProps {}

function GlobalStyle(props: GlobalStyleProps) {
  const { theme } = useColorThemeValue();
  return <Global styles={globalStyle(theme)} />;
}

const globalStyle = (theme: ColorTheme) => css`
  html,
  body,
  #root {
    margin: 0;
    padding: 0;
    height: 100%;
    background: ${themeColor.body[theme]};
  }
  header {
    background: ${themeColor.header[theme]};
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

  h1,
  h2,
  h3,
  h4,
  h5,
  h6,
  p,
  span,
  a {
    color: ${themeColor.font[theme]};
  }
  a {
    text-decoration: none;
    cursor: pointer;
  }
`;

export default GlobalStyle;
