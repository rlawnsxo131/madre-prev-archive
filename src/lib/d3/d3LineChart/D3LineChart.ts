import { axisBottom, axisLeft, line, scaleLinear } from 'd3';
import D3Common from '../d3Common';
import {
  D3Data,
  D3DoubleNumberArray,
  D3Path,
  D3Selection,
} from '../d3Common/d3CommonTypes';
import {
  D3LineChartConstructorParams,
  D3LineChartSetAxisParams,
  D3LineChartSetLineParams,
} from './d3LineChartTypes';

export default class D3LineChart extends D3Common {
  private svg: D3Selection;
  private width: number;
  private height: number;
  private xDomain: D3DoubleNumberArray;
  private yDomain: D3DoubleNumberArray;
  private xRange: D3DoubleNumberArray;
  private yRange: D3DoubleNumberArray;
  private data: D3Data;
  private path: D3Path | null = null;
  private pathLength?: number = 0;
  private axisFontSizeTimesThree = 24;

  constructor({
    container,
    width,
    height,
    className = '',
    xDomain,
    yDomain,
    xRange,
    yRange,
    data,
  }: D3LineChartConstructorParams) {
    super();
    this.svg = this.appendSvg({
      container,
      width,
      height,
      className,
    });
    this.width = width;
    this.height = height;
    this.xDomain = this.getExtent(xDomain);
    this.yDomain = this.getExtent(yDomain);
    this.xRange = xRange;
    this.yRange = yRange;
    this.data = data;
  }

  private xScale() {
    return scaleLinear().domain(this.xDomain).range(this.xRange);
  }

  private yScale() {
    return scaleLinear().domain(this.yDomain).range(this.yRange);
  }

  setAxis({
    xTicks = 0,
    yTicks = 0,
    xTickSize = 0,
    yTickSize = 0,
    xClass = '',
    yClass = '',
    axisFontSize,
  }: D3LineChartSetAxisParams) {
    if (axisFontSize) {
      this.axisFontSizeTimesThree = axisFontSize * 3;
    }
    const translate = this.axisFontSizeTimesThree;
    const svg = this.svg;

    const xAxisSvg = svg
      .append('g')
      .attr('class', xClass)
      .attr(
        'transform',
        `translate(${translate}, ${this.height - yTicks - yTickSize})`,
      );

    const yAxisSvg = svg
      .append('g')
      .attr('class', yClass)
      .attr('transform', `translate(${translate}, ${translate})`);

    const xAxis = axisBottom(this.xScale()).tickSize(xTickSize).ticks(xTicks);
    const yAxis = axisLeft(this.yScale()).tickSize(yTickSize).ticks(yTicks);

    xAxis(xAxisSvg);
    yAxis(yAxisSvg);
  }

  setLine({ color = 'black', strokeWidth = 1 }: D3LineChartSetLineParams) {
    const translate = this.axisFontSizeTimesThree;
    const xScale = this.xScale();
    const yScale = this.yScale();
    const linearGenerator = line()
      .x((d) => xScale(d[0]))
      .y((d) => yScale(d[1]));

    const path = this.svg
      .append('path')
      .attr('fill', 'none')
      .attr('d', `${linearGenerator(this.data)}`)
      .attr('stroke-width', strokeWidth)
      .attr('stroke', color)
      // .attr('transform', `translate(${translate}, ${translate})`);
      .attr('transform', `translate(${translate + strokeWidth}, ${translate})`);

    this.path = path;
    this.pathLength = path.node()?.getTotalLength() ?? 0;
  }
}
