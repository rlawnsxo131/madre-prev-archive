import { css } from '@emotion/react';
import HomeSectionGraph from './HomeSectionGraph';
import HomeSectionThinkAbout from './HomeSectionThinkAbout';

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

HomeSection.ThinkAbout = HomeSectionThinkAbout;
HomeSection.Graph = HomeSectionGraph;

export default HomeSection;
