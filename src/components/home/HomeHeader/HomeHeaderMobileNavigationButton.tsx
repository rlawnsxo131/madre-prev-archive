import { css } from '@emotion/react';
import { MenuIcon } from '../../../image/icons';
import { palette, themeColor } from '../../../styles';
import useSetHomeMobileNavigation from '../../../hooks/home/useSetHomeMobileNavigation';

interface HomeHeaderMobileNavigationButtonProps {}

function HomeHeaderMobileNavigationButton(
  props: HomeHeaderMobileNavigationButtonProps,
) {
  const onClick = useSetHomeMobileNavigation();

  return (
    <button css={block} onClick={onClick}>
      <MenuIcon />
    </button>
  );
}

const block = css`
  background: inherit;
  border: none;
  box-shadow: none;
  border-radius: 0;
  overflow: visible;
  cursor: pointer;
  display: flex;
  align-items: center;
  padding: 0.5rem 0;
  border-radius: 3px;
  color: ${palette.gray['500']};
  svg {
    width: 1.125rem;
    height: 1.125rem;
    fill: ${themeColor.fill['light']};
  }
  &:hover {
    svg {
      opacity: 0.5;
    }
  }
`;

export default HomeHeaderMobileNavigationButton;
