import { css } from '@emotion/react';
import { useEffect, useRef } from 'react';
import D3LineChart from '../../../lib/d3/D3LineChart';
import {
  D3LineChartKindsType,
  LineAndPointInitializeParams,
} from '../../../lib/d3/types';

interface LineChartProps extends Omit<LineAndPointInitializeParams, 'svg'> {
  width: number;
  height: number;
  colors: string[];
  xKeys: string[];
  yKeys: string[];
  kinds: D3LineChartKindsType[];
}

function LineChart(params: LineChartProps) {
  const svgRef = useRef<SVGSVGElement | null>(null);

  useEffect(() => {
    if (!svgRef.current || !params) return;
    const chart = new D3LineChart();
    const {
      data,
      xDomainData,
      yDomainData,
      xRange,
      yRange,
      xTickSize,
      xTicks,
      yTickSize,
      yTicks,
      xTransform,
      yTransform,
      colors,
      xKeys,
      yKeys,
      kinds,
    } = params;
    chart.initialize({
      svg: svgRef.current,
      data,
      xDomainData,
      yDomainData,
      xRange,
      yRange,
      xTickSize,
      xTicks,
      yTickSize,
      yTicks,
      xTransform,
      yTransform,
    });
    for (let i = 0; i < xKeys.length; i++) {
      const xKey = xKeys[i];
      const yKey = yKeys[i];
      const color = colors[i];
      if (kinds.includes('circle')) {
        chart.setCircle({
          xKey,
          yKey,
          color,
        });
      }
      if (kinds.includes('line')) {
        chart.setLine({
          xKey,
          yKey,
          color,
        });
        chart.strokeAnimate();
        if (kinds.includes('area')) {
          chart.setArea({
            xKey,
            yKey,
            color,
          });
          chart.areaAnimate();
        }
      }
    }
  }, [svgRef.current, params]);

  return <svg css={block(params.width, params.height)} ref={svgRef} />;
}

const block = (width: number, height: number) => css`
  width: ${width}px;
  height: ${height}px;
`;

export default LineChart;
