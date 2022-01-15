import { memo } from 'react';
import { css } from '@emotion/react';
import { themeColor, transitions, zIndexes } from '../../../styles';
import HomeNavigationLink from './HomeNavigationLink';
import HomeTemplateStyles from './HomeTemplate.styles';

interface HomeNavigationLinksProps {
  visible: boolean;
}

function HomeNavigationLinks({ visible }: HomeNavigationLinksProps) {
  return (
    <nav css={block(visible)}>
      <ul css={HomeTemplateStyles.listBlock}>
        <HomeNavigationLink to="/" text="홈" />
        <HomeNavigationLink to="/preview" text="미리보기" />
        <HomeNavigationLink to="/guides" text="가이드 및 튜토리얼" />
        <HomeNavigationLink to="/notice" text="공지사항" />
        <HomeNavigationLink to="/policy" text="서비스 정책" />
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

export default memo(HomeNavigationLinks);
