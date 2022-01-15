import { memo } from 'react';
import { css } from '@emotion/react';
import AppNavigationLink from './PublicNavigationLink';
import AppNavigationStyles from './PublicNavigation.styles';
import { themeColor, transitions, zIndexes } from '../../../styles';

interface AppNavigationLinksProps {
  visible: boolean;
}

function AppNavigationLinks({ visible }: AppNavigationLinksProps) {
  return (
    <nav css={block(visible)}>
      <ul css={AppNavigationStyles.listBlock}>
        <AppNavigationLink to="/" text="홈" />
        <AppNavigationLink to="/preview" text="미리보기" />
        <AppNavigationLink to="/guides" text="가이드 및 튜토리얼" />
        <AppNavigationLink to="/notice" text="공지사항" />
        <AppNavigationLink to="/policy" text="서비스 정책" />
      </ul>
    </nav>
  );
}

const block = (visible: boolean) => css`
  position: absolute;
  top: 3.25rem;
  left: 0.5rem;
  width: 15.15rem;
  height: auto;
  padding: 0.25rem 0.5rem;
  display: flex;
  flex-direction: column;
  z-index: ${zIndexes.dropdownItem};
  border-radius: 0.25rem;
  background: ${themeColor.navigation['light']};
  box-shadow: ${themeColor.shadow['light']};
  transform-origin: top;
  ${visible
    ? css`
        animation: ${transitions.scaleUp} 0.125s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.scaleDown} 0.125s forwards ease-in-out;
      `};
`;

export default memo(AppNavigationLinks);
