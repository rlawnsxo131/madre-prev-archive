import { css } from '@emotion/react';
import Button from '../../common/Button';
import { GoogleIcon, MenuIcon } from '../../../image/icons';
import HomeNavigation from './HomeNavigation';
import { useHomeValue, useHomeActions } from '../../../atoms/homeState';
import { palette, themeColor } from '../../../styles';

interface HomeHeaderItemsProps {}

function HomeHeaderItems(props: HomeHeaderItemsProps) {
  const { visible } = useHomeValue();
  const { handleNavigation } = useHomeActions();

  return (
    <div css={block}>
      <div css={itemWrapper}>
        <Button size="medium" icon={<GoogleIcon />} outline>
          Sign in with Google
        </Button>
      </div>
      <div css={itemWrapper}>
        <button css={button} onClick={handleNavigation}>
          <MenuIcon />
        </button>
        <HomeNavigation visible={visible} />
      </div>
    </div>
  );
}

const block = css`
  position: relative;
  display: flex;
  align-items: center;
`;

const itemWrapper = css`
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

const button = css`
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
    fill: ${themeColor.fill['light']};
  }
`;

export default HomeHeaderItems;
