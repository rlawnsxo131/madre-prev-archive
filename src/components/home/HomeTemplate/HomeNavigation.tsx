import { css } from '@emotion/react';
import Button from '../../common/Button';
import { ThemeTrigger } from '../../common/Theme';
import { GoogleIcon, MenuIcon } from '../../../image/icons';
import HomeNavigationItem from './HomeNavigationItem';
import { useHomeAction, useHomeState } from '../../../atoms/homeState';
import { DarkmodeThemeType } from '../../../atoms/darkmodeState';
import palette, { themeColor } from '../../../styles/palette';

interface HomeNavigationProps {}

function HomeNavigation(props: HomeNavigationProps) {
  const { theme, visible } = useHomeState();
  const { handleNavigation } = useHomeAction();

  return (
    <nav css={block}>
      <ul>
        <li>
          <Button size="medium" icon={<GoogleIcon />} outline>
            Sign in with Google
          </Button>
        </li>
        <li>
          <button css={buttonStyle(theme)} onClick={handleNavigation}>
            <MenuIcon />
          </button>
          <HomeNavigationItem theme={theme} visible={visible} />
        </li>
        <li>
          <ThemeTrigger />
        </li>
      </ul>
    </nav>
  );
}

const block = css`
  position: relative;
  display: flex;
  align-items: center;
  ul,
  li {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  li {
    &:not(:nth-of-type(3)) {
      padding: 0 0.5rem;
    }
  }
`;

const buttonStyle = (theme: DarkmodeThemeType) => css`
  background: inherit;
  border: none;
  box-shadow: none;
  border-radius: 0;
  overflow: visible;
  cursor: pointer;
  display: flex;
  align-items: center;
  padding: 0.5rem;
  border-radius: 3px;
  color: ${palette.gray['500']};
  svg {
    width: 1.125rem;
    height: 1.125rem;
    fill: ${themeColor.fill[theme]};
  }
`;

export default HomeNavigation;
