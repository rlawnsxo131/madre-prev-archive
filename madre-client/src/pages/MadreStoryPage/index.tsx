import { css } from '@emotion/react';
import MadreStorySectionGraph from './components/MadreStorySectionGraph';
import MadreStorySectionThinkAbout from './components/MadreStorySectionThinkAbout';

interface indexProps {}

function MadreStoryPage(props: indexProps) {
  return (
    <div css={block}>
      <MadreStorySectionThinkAbout />
      <MadreStorySectionGraph />
    </div>
  );
}

const block = css`
  display: flex;
  flex-direction: column;
`;

export default MadreStoryPage;
