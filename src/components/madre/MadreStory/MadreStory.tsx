import { css } from '@emotion/react';
import MadreStorySectionGraph from './MadreStorySectionGraph';
import MadreStorySectionThinkAbout from './MadreStorySectionThinkAbout';

interface MadreStoryProps {}

function MadreStory(props: MadreStoryProps) {
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

export default MadreStory;
