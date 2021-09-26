import {
  axisBottom,
  axisLeft,
  axisRight,
  axisTop,
  curveBasis,
  easeSinInOut,
  line,
  scaleLinear,
} from 'd3';
import D3Common from '../d3Common';
import {
  D3DoubleNumberArray,
  D3Path,
  D3Selection,
} from '../d3Common/d3CommonTypes';
import {
  D3AxisChartConstructorParams,
  D3AxisChartSetAxisParams,
  D3AxisChartSetLineParams,
} from './d3AxisChartTypes';

/**
 * strockWidth: linechart strock width
 * axisMaxUnitExpressionLength: maximum number of characters displayed
 */
export default class D3AxisChart extends D3Common {
  private svg: D3Selection;
  private width: number;
  private height: number;
  private xDomain: D3DoubleNumberArray;
  private yDomain: D3DoubleNumberArray;
  private xRange: D3DoubleNumberArray;
  private yRange: D3DoubleNumberArray;
  // axis attribute
  private axisFontSize = 10;
  private axisMaxUnitExpressionLength = 0;
  // straight attribute
  private strockWidth = 1;

  constructor({
    container,
    width,
    height,
    className = '',
    xDomain,
    yDomain,
    xRange,
    yRange,
  }: D3AxisChartConstructorParams) {
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
    return scaleLinear().domain(this.xDomain).nice().range(this.xRange);
  }

  private yScale() {
    return scaleLinear().domain(this.yDomain).nice().range(this.yRange);
  }

  setAxis({
    xTicks = 0,
    yTicks = 0,
    xTickSize = 0,
    yTickSize = 0,
    xClass = '',
    yClass = '',
    axisFontSize = 10,
    axisMaxUnitExpressionLength = 0,
    xTickFormat = (d, i) => `${d}`,
    yTickFormat = (d, i) => `${d}`,
    xGridClass = '',
    yGridClass = '',
  }: D3AxisChartSetAxisParams) {
    this.axisFontSize = axisFontSize;
    this.axisMaxUnitExpressionLength = axisMaxUnitExpressionLength;

    const xAxisSvg = this.svg
      .append('g')
      .attr('class', xClass)
      .attr(
        'transform',
        `translate(
          ${this.axisMaxUnitExpressionLength}, 
          ${this.height - yTicks - yTickSize})`,
      )
      .style('font-size', this.axisFontSize);

    const yAxisSvg = this.svg
      .append('g')
      .attr('class', yClass)
      .attr(
        'transform',
        `translate(
          ${this.axisMaxUnitExpressionLength},
          ${this.axisMaxUnitExpressionLength})`,
      )
      .style('font-size', this.axisFontSize);

    const xAxis = axisBottom(this.xScale())
      .tickSize(xTickSize)
      .ticks(xTicks)
      .tickFormat(xTickFormat);

    const yAxis = axisLeft(this.yScale())
      .tickSize(yTickSize)
      .ticks(yTicks)
      .tickFormat(yTickFormat);

    xAxis(xAxisSvg);
    yAxis(yAxisSvg);

    // grid line
    const xGridSvg = this.svg
      .append('g')
      .attr('class', xGridClass)
      .attr(
        'transform',
        `translate(
          ${this.axisMaxUnitExpressionLength},
          ${this.axisMaxUnitExpressionLength})`,
      )
      .style('font-size', this.axisFontSize);

    const yGirdSvg = this.svg
      .append('g')
      .attr('class', yGridClass)
      .attr(
        'transform',
        `translate(
          ${this.axisMaxUnitExpressionLength},
          ${this.height - yTicks - yTickSize}
        )`,
      )
      .style('fornt-size', this.axisFontSize);

    const xGrid = axisRight(this.yScale())
      .tickSize(this.width)
      .ticks(5)
      .tickFormat(() => '');

    const yGrid = axisTop(this.xScale())
      .tickSize(this.height)
      .ticks(5)
      .tickFormat(() => '');

    xGrid(xGridSvg);
    yGrid(yGirdSvg);
  }

  // axis background grid line
  setGrid() {}

  // line graph
  setLine({
    data,
    color = 'black',
    strokeWidth = 1,
    lineType = 'STRAIGHT',
    animate = false,
  }: D3AxisChartSetLineParams) {
    this.strockWidth = strokeWidth;

    const xScale = this.xScale();
    const yScale = this.yScale();

    let linearGenerator = line()
      .x((d) => xScale(d[0]))
      .y((d) => yScale(d[1]));

    if (lineType === 'CURVE') {
      linearGenerator.curve(curveBasis);
    }

    const path = this.svg
      .append('path')
      .attr('fill', 'none')
      .attr('d', `${linearGenerator(data)}`)
      .attr('stroke-width', strokeWidth)
      .attr('stroke', color)
      .attr(
        'transform',
        `translate(${this.axisMaxUnitExpressionLength + strokeWidth / 2}, 
          ${this.axisMaxUnitExpressionLength})
        `,
      );

    const pathLength = path.node()?.getTotalLength();

    if (animate && path && pathLength) {
      this.lineAnimate(path, pathLength);
    }
  }

  private lineAnimate(path: D3Path, pathLength: number) {
    path
      .attr('stroke-dashoffset', pathLength)
      .attr('stroke-dasharray', pathLength)
      .transition()
      .ease(easeSinInOut)
      .duration(1500)
      .attr('stroke-dashoffset', 0); //시작점
  }
}
