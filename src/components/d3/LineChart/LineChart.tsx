import { useRef } from 'react';

interface LineChartProps {}

function LineChart(params: LineChartProps) {
  const svgRef = useRef<SVGSVGElement | null>(null);

  return <svg ref={svgRef} />;
}

export default LineChart;
