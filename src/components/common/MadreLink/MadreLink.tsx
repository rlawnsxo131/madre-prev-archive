import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import { themePalette } from '../../../styles';

type ParentsDirection = 'row' | 'column';

interface MadreLinkProps {
  to: string;
  displayName: string;
  parentDirection?: ParentsDirection;
}

function MadreLink({
  to,
  displayName,
  parentDirection = 'row',
}: MadreLinkProps) {
  return (
    <NavLink css={link(parentDirection)} to={to}>
      {displayName}
    </NavLink>
  );
}

const link = (parentDirection: ParentsDirection) => css`
  display: flex;
  &.active {
    color: ${themePalette.anchor_active1};
  }

  ${parentDirection === 'row' &&
  css`
    display: flex;
    align-items: center;
    font-size: 0.875rem;
    font-weight: bold;
    padding: 0.5rem 1rem;
  `}

  ${parentDirection === 'column' &&
  css`
    flex-flow: row wrap;
    align-items: center;
    font-size: 0.9rem;
    font-weight: bold;
    padding: 0.5rem 0.25rem;
  `}
`;

export default MadreLink;
