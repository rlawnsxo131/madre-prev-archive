import { css } from '@emotion/react';
import { DarkmodeThemeType } from '../../atoms/darkmodeState';
import { GoogleIcon } from '../../image/icons';
import DropArrowIcon from '../../image/icons/DropArrowIcon';
import palette, { themeColor } from '../../styles/palette';
import Button from '../common/Button';
import { ThemeTrigger } from '../common/Theme';
import HomeDropDownItem from './HomeDropDownItem';
import useHome from './hooks/useHome';

interface HomeNavigationProps {}

function HomeNavigation(props: HomeNavigationProps) {
  const { theme, visible, handleVisible } = useHome();

  return (
    <nav css={block}>
      <ul>
        <li>
          <button css={buttonStyle(theme)} onClick={handleVisible}>
            <p>자료</p>
            <DropArrowIcon />
          </button>
          <HomeDropDownItem theme={theme} visible={visible} />
        </li>
        <li>
          <Button size="medium" icon={<GoogleIcon />} outline>
            Sign in with Google
          </Button>
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
      margin-right: 1rem;
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
  color: ${palette.gray['5']};
  p {
    margin: 0;
    padding: 0;
    font-size: 0.9375rem;
    font-weight: 500;
  }
  svg {
    width: 10px;
    height: 10px;
    margin-left: 0.4rem;
    fill: ${themeColor.fill[theme]};
  }
`;

export default HomeNavigation;
