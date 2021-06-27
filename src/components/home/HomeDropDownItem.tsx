import { css } from '@emotion/react';
import { memo } from 'react';
import { NavLink } from 'react-router-dom';
import { DarkmodeThemeType } from '../../atoms/darkmodeState';
import { themeColor } from '../../styles/palette';
import zIndexes from '../../styles/zIndexes';

interface HomeDropDownItemProps {
  theme: DarkmodeThemeType;
  visible: boolean;
}

function HomeDropDownItem({ theme, visible }: HomeDropDownItemProps) {
  return (
    <div css={block(theme, visible)}>
      <NavLink css={link} to="/notice">
        공지사항
      </NavLink>
      <NavLink css={link} to="/guides">
        가이드 및 튜토리얼
      </NavLink>
      <NavLink css={link} to="/policy">
        서비스 정책
      </NavLink>
    </div>
  );
}

const block = (theme: DarkmodeThemeType, visible: boolean) => css`
  position: absolute;
  top: 3.25rem;
  left: 0;
  width: 18rem;
  height: auto;
  padding: 0.25rem 0.5rem;
  display: flex;
  flex-direction: column;
  z-index: ${zIndexes.dropdownItem};
  border-radius: 0.25rem;
  background: ${themeColor.commonDeepDark[theme]};
  box-shadow: ${themeColor.shadow[theme]};
  opacity: ${visible ? 1 : 0};
`;

const link = css`
  display: flex;
  flex-flow: row wrap;
  align-items: center;
  font-size: 1rem;
  padding: 0.5rem 0.25rem 0.5rem 0.25rem;
  &.active {
    font-weight: bold;
  }
`;

export default memo(HomeDropDownItem);
