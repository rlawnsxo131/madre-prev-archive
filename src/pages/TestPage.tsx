import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import { D3LineChart } from '../lib/d3';
import { D3Data } from '../lib/d3/types/d3CommonTypes';
import { getRandomIntInclusive } from '../lib/utils';

interface TestPageProps {}

function TestPage(props: TestPageProps) {
  const ref = useRef<HTMLDivElement | null>(null);
  useEffect(() => {
    if (!ref.current) return;
    const min = 40;
    const max = 87;
    const data: D3Data = Array.from({ length: 50 }).map((_, i) => [
      i * 2,
      getRandomIntInclusive(min, max),
    ]);
    const xDomainData = data.map((v) => v[0]);
    const yDomainData = data.map((v) => v[1]);
    console.log('data: ', data);
    console.log('xDomainData: ', xDomainData);
    console.log('yDomainData: ', yDomainData);
    const width = 500;
    const height = 320;

    const chart = new D3LineChart({
      container: ref.current,
      width,
      height,
      xDomainData,
      yDomainData,
      xRange: [0, 500],
      yRange: [0, 320],
    });
    console.log(chart);
  }, [ref.current]);
  return <div css={block} ref={ref}></div>;
}

const block = css`
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
`;

export default TestPage;
