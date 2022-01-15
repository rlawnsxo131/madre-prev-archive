import { memo } from 'react';
import { css } from '@emotion/react';
import PublicNavigationLink from './PublicNavigationLink';
import PublicNavigationStyles from './PublicNavigation.styles';
import { themeColor, transitions, zIndexes } from '../../../styles';
import useTransitionTimeoutEffect from '../../../lib/hooks/useTransitionTimeoutEffect';

interface PublicNavigationProps {
  visible: boolean;
}

function PublicNavigation({ visible }: PublicNavigationProps) {
  const closed = useTransitionTimeoutEffect({ visible });

  if (!visible && closed) return null;

  return (
    <nav css={block(visible)}>
      <ul css={PublicNavigationStyles.listBlock}>
        <PublicNavigationLink to="/" text="홈" />
        <PublicNavigationLink to="/preview" text="미리보기" />
        <PublicNavigationLink to="/guides" text="가이드 및 튜토리얼" />
        <PublicNavigationLink to="/notice" text="공지사항" />
        <PublicNavigationLink to="/policy" text="서비스 정책" />
      </ul>
    </nav>
  );
}

const block = (visible: boolean) => css`
  position: absolute;
  top: 3.25rem;
  left: calc((12rem - (1.125rem * 2)) * -1);
  width: 12rem;
  height: auto;
  padding: 0.25rem 0.5rem;
  display: flex;
  flex-direction: column;
  z-index: ${zIndexes.dropdownItem};
  border-radius: 0.25rem;
  background: ${themeColor.navigation['light']};
  box-shadow: ${themeColor.shadow['light']};
  transform-origin: top right;
  ${visible
    ? css`
        animation: ${transitions.scaleUp} 0.25s forwards ease-in-out;
      `
    : css`
        animation: ${transitions.scaleDown} 0.25s forwards ease-in-out;
      `};
`;

export default memo(PublicNavigation);
