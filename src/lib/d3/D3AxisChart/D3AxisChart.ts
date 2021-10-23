import {
  axisBottom,
  axisLeft,
  axisRight,
  axisTop,
  curveBasis,
  easeSinInOut,
  extent,
  line,
  scaleLinear,
  select,
} from 'd3';
import D3Common from '../D3Common';
import {
  D3Axis,
  D3DoubleNumberArray,
  D3Line,
  D3Selection,
  D3SelectionSVG,
} from '../D3Common/D3CommonTypes';
import {
  D3AxisChartConstructorParams,
  D3AxisChartDrawLineParams,
  D3AxisChartSetAxisBackgroundGridParams,
  D3AxisChartSetAxisParams,
} from './D3AxisChartTypes';

export default class D3AxisChart extends D3Common {
  /**
   * @Initialize
   */
  private svg: D3SelectionSVG;
  private width: number;
  private height: number;
  private xDomain: D3DoubleNumberArray;
  private yDomain: D3DoubleNumberArray;
  private xRange: D3DoubleNumberArray;
  private yRange: D3DoubleNumberArray;

  /**
   * @Axis
   * axisMaxUnitExpressionLength: maximum number of characters displayed
   * axisX: axis draw function
   * axisY: axis draw function
   */
  private axisFontSize = 10;
  private axisMaxUnitExpressionLength = 0;
  private axisXTicks = 0;
  private axisXTickSize = 0;
  private axisYTicks = 0;
  private axisYTickSize = 0;
  private axisXSvg: D3Selection = null;
  private axisYSvg: D3Selection = null;
  private axisX: D3Axis = null;
  private axisY: D3Axis = null;

  /**
   * @AxisGrid
   */
  private axisGridXSvg: D3Selection = null;
  private axisGridYSvg: D3Selection = null;
  private axisGridX: D3Axis = null;
  private axisGridY: D3Axis = null;

  /**
   * @LineGraph
   */
  private lineGenerator: D3Line = null;

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
    this.xDomain = this.serializedExtent(extent(xDomain));
    this.yDomain = this.serializedExtent(extent(yDomain));
    this.xRange = xRange;
    this.yRange = yRange;
  }

  private xScale() {
    return scaleLinear().domain(this.xDomain).range(this.xRange).nice();
  }

  private yScale() {
    return scaleLinear().domain(this.yDomain).range(this.yRange).nice();
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
  }: D3AxisChartSetAxisParams) {
    this.axisFontSize = axisFontSize;
    this.axisMaxUnitExpressionLength = axisMaxUnitExpressionLength;
    this.axisXTicks = xTicks;
    this.axisXTickSize = xTickSize;
    this.axisYTicks = yTicks;
    this.axisYTickSize = yTickSize;

    this.axisXSvg = this.svg
      .append('g')
      .attr('class', xClass)
      .attr(
        'transform',
        `translate(
          ${this.axisMaxUnitExpressionLength - this.axisXTickSize}, 
          ${this.height - this.axisYTicks - this.axisYTickSize}
        )`,
      )
      .style('font-size', this.axisFontSize);

    this.axisYSvg = this.svg
      .append('g')
      .attr('class', yClass)
      .attr(
        'transform',
        `translate(
          ${this.axisMaxUnitExpressionLength - this.axisXTickSize},
          ${this.axisMaxUnitExpressionLength}
        )`,
      )
      .style('font-size', this.axisFontSize);

    this.axisX = axisBottom(this.xScale())
      .tickSize(this.axisXTickSize)
      .ticks(this.axisXTicks)
      .tickFormat(xTickFormat);

    this.axisY = axisLeft(this.yScale())
      .tickSize(this.axisYTickSize)
      .ticks(this.axisYTicks)
      .tickFormat(yTickFormat);
  }

  setAxisBackgroundGrid({
    direction,
    xClass = '',
    yClass = '',
    xTickFormat = () => '',
    yTickFormat = () => '',
  }: D3AxisChartSetAxisBackgroundGridParams) {
    if (direction.x) {
      this.axisGridXSvg = this.svg
        .append('g')
        .attr('class', xClass)
        .attr(
          'transform',
          `translate(
        ${this.axisMaxUnitExpressionLength - this.axisXTickSize},
        ${this.axisMaxUnitExpressionLength})`,
        )
        .style('font-size', this.axisFontSize);

      this.axisGridX = axisRight(this.yScale())
        .tickSize(this.width)
        .ticks(5)
        .tickFormat(xTickFormat);
    }

    if (direction.y) {
      this.axisGridYSvg = this.svg
        .append('g')
        .attr('class', yClass)
        .attr(
          'transform',
          `translate(
          ${this.axisMaxUnitExpressionLength - this.axisXTickSize},
          ${this.height - this.axisYTicks - this.axisYTickSize}
        )`,
        )
        .style('fornt-size', this.axisFontSize);

      this.axisGridY = axisTop(this.xScale())
        .tickSize(this.height)
        .ticks(5)
        .tickFormat(yTickFormat);
    }
  }

  setArea() {}

  drawAxis() {
    if (this.axisXSvg && this.axisX) {
      this.axisX(this.axisXSvg);
    }
    if (this.axisYSvg && this.axisY) {
      this.axisY(this.axisYSvg);
    }
  }

  drawGrid() {
    if (this.axisGridXSvg && this.axisGridX) {
      this.axisGridX(this.axisGridXSvg);
    }
    if (this.axisGridYSvg && this.axisGridY) {
      this.axisGridY(this.axisGridYSvg);
    }
  }

  drawLine({
    data,
    color = 'black',
    strokeWidth = 2,
    lineType = 'STRAIGHT',
    animate = false,
  }: D3AxisChartDrawLineParams) {
    const xScale = this.xScale();
    const yScale = this.yScale();

    this.lineGenerator = line()
      .x((d) => xScale(d[0]))
      .y((d) => yScale(d[1]));

    if (lineType === 'CURVE') {
      this.lineGenerator.curve(curveBasis);
    }

    const path = this.svg
      .append('path')
      .attr('fill', 'none')
      .attr('d', `${this.lineGenerator(data.d3Position)}`)
      .attr('stroke-width', strokeWidth)
      .attr('stroke', color)
      .attr(
        'transform',
        `translate(
          ${
            this.axisMaxUnitExpressionLength -
            this.axisXTickSize +
            strokeWidth / 2
          },
          ${this.axisMaxUnitExpressionLength}
        )
        `,
      )
      .on('mouseover', function (d) {
        select(this)
          .style('stroke-width', strokeWidth * 2.5)
          .style('cursor', 'pointer');
      })
      .on('mouseout', function (d) {
        select(this)
          .style('stroke-width', strokeWidth)
          .style('cursor', 'default');
      });

    const pathLength = path.node()?.getTotalLength();

    if (animate && path && pathLength) {
      path
        .attr('stroke-dashoffset', pathLength)
        .attr('stroke-dasharray', pathLength)
        .transition()
        .ease(easeSinInOut)
        .duration(1500)
        .attr('stroke-dashoffset', 0); //시작점
    }
  }
}
