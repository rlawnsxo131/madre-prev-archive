import { css } from '@emotion/react';

interface HomeSectionProps {
  children: React.ReactNode;
}

function HomeSection({ children }: HomeSectionProps) {
  return <div css={block}>{children}</div>;
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default HomeSection;
