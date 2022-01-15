import { css } from '@emotion/react';
import { NavLink } from 'react-router-dom';
import HomeTemplateStyles from './HomeTemplate.styles';

interface HomeNavigationLinkProps {
  text: string;
  to: string;
}

function HomeNavigationLink({ text, to }: HomeNavigationLinkProps) {
  return (
    <li css={HomeTemplateStyles.listBlock}>
      <NavLink css={link} to={to}>
        {text}
      </NavLink>
    </li>
  );
}

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

export default HomeNavigationLink;
