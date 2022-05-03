import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import { standardColor } from '../../../styles';

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
  const className = ({ isActive }: { isActive: boolean }) =>
    isActive ? 'active' : undefined;

  return (
    <NavLink css={link(parentDirection)} to={to} className={className}>
      {displayName}
    </NavLink>
  );
}

const link = (parentDirection: ParentsDirection) => css`
  display: flex;
  &.active {
    color: ${standardColor.navigation.active};
  }

  ${parentDirection === 'row'
    ? css`
        display: flex;
        align-items: center;
        font-size: 0.875rem;
        font-weight: bold;
        padding: 0 1rem;
      `
    : css`
        flex-flow: row wrap;
        align-items: center;
        font-size: 0.9rem;
        font-weight: bold;
        padding: 0.5rem 0.25rem 0.5rem 0.25rem;
      `}
`;

export default MadreLink;
