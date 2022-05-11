import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import { themePalette } from '../../../styles';

interface MadreImageLinkProps {
  children: React.ReactNode;
  to: string;
  onClick?: (e: React.MouseEvent<HTMLAnchorElement>) => void;
}

function MadreImageLink({ children, to, onClick }: MadreImageLinkProps) {
  return (
    <NavLink css={link} to={to} onClick={onClick}>
      {children}
    </NavLink>
  );
}

const link = css`
  display: flex;
  justify-content: center;
  align-items: center;
  &.active {
    svg {
      fill: ${themePalette.anchor_active1};
    }
  }
`;

export default MadreImageLink;
