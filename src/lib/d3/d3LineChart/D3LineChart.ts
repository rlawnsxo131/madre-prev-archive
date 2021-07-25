import { axisBottom, axisLeft, csv, scaleLinear } from 'd3';
import D3Common from '../d3Common';
import { D3DoubleNumberArray, D3Selection } from '../d3Common/d3CommonTypes';
import {
  D3LineChartConstructorParams,
  D3LineChartSetAxisParams,
} from './d3LineChartTypes';

export default class D3LineChart extends D3Common {
  private svg: D3Selection;
  private width: number;
  private height: number;
  private xDomain: D3DoubleNumberArray;
  private yDomain: D3DoubleNumberArray;
  private xRange: D3DoubleNumberArray;
  private yRange: D3DoubleNumberArray;

  constructor({
    container,
    width,
    height,
    className = '',
    xDomain,
    yDomain,
    xRange,
    yRange,
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
    xClassName = '',
    yClassName = '',
  }: D3LineChartSetAxisParams) {
    const svg = this.svg;

    const xAxisSvg = svg
      .append('g')
      .attr('class', xClassName)
      .attr('transform', `translate(24, ${this.height - 10 - 6})`);

    const yAxisSvg = svg
      .append('g')
      .attr('className', yClassName)
      .attr('transform', `translate(24, 24)`);

    const xAxis = axisBottom(this.xScale()).tickSize(xTickSize).ticks(xTicks);
    const yAxis = axisLeft(this.yScale()).tickSize(yTickSize).ticks(yTicks);

    xAxis(xAxisSvg);
    yAxis(yAxisSvg);
  }
}
