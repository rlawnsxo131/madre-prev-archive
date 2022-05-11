import { css } from '@emotion/react';
import { useMemo } from 'react';
import useMatchedRoute from '../../../hooks/useMatchedRoute';
import { basicStyles, themePalette } from '../../../styles';

interface MadreButtonLinkProps {
  children: React.ReactNode;
  onClick: (e?: React.MouseEvent<HTMLButtonElement>) => void;
  matchPath?: string;
}

function MadreButtonLink({
  children,
  onClick,
  matchPath,
}: MadreButtonLinkProps) {
  const route = useMatchedRoute();
  const isActive = useMemo(() => {
    if (!matchPath) return false;
    return route?.startsWith(matchPath);
  }, [route, matchPath]);

  return (
    <button css={[basicStyles.button, button(isActive)]} onClick={onClick}>
      {children}
    </button>
  );
}

const button = (isActive?: boolean) => css`
  display: flex;
  justify-content: center;
  align-items: center;
  ${isActive &&
  css`
    color: ${themePalette.anchor_active1};
    svg {
      fill: ${themePalette.anchor_active1};
    }
  `}
`;

export default MadreButtonLink;
