import {
  area,
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
  selectAll,
} from 'd3';
import D3Common from '../D3Common';
import {
  D3Axis,
  D3DoubleNumberArray,
  D3Selection,
  D3SelectionSVG,
} from '../D3Common/D3CommonTypes';
import {
  D3AxisChartConstructorParams,
  D3AxisChartDrawAreaParams,
  D3AxisChartDrawLineParams,
  D3AxisChartLineType,
  D3AxisChartSetAxisBackgroundGridParams,
  D3AxisChartSetAxisParams,
  D3Margin,
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
  private margin: D3Margin;

  /**
   * @Axis
   * axisX: axis draw function
   * axisY: axis draw function
   */
  private axisFontSize = 10;
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
  private axisGridXTicks: number = 5;
  private axisGridYTicks: number = 5;

  /**
   * @AxisLine @AxisArea
   */
  private lineKey = 'line-';
  private areaKey = 'area-';
  private strokeWidth: number = 2;
  private lineType: D3AxisChartLineType = 'STRAIGHT';

  constructor({
    container,
    width,
    height,
    className = '',
    xDomain,
    yDomain,
    margin,
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
    this.margin = margin;
    this.xRange = [0, width - (margin.left + margin.right)];
    this.yRange = [height - (margin.top + margin.bottom), 0];
  }

  private xScale() {
    return scaleLinear().domain(this.xDomain).range(this.xRange).nice();
  }

  private yScale() {
    return scaleLinear().domain(this.yDomain).range(this.yRange).nice();
  }

  private onMouseOverAction(targetClass?: string) {}

  private onMouseOutAction(targetClass?: string) {}

  setAxis({
    xTicks = 0,
    yTicks = 0,
    xTickSize = 0,
    yTickSize = 0,
    xClass = '',
    yClass = '',
    axisFontSize = 10,
    xTickFormat = (d, i) => `${d}`,
    yTickFormat = (d, i) => `${d}`,
  }: D3AxisChartSetAxisParams) {
    this.axisFontSize = axisFontSize;
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
          ${this.margin.left}, 
          ${this.height - this.margin.top}
        )`,
      )
      .style('font-size', this.axisFontSize);

    this.axisYSvg = this.svg
      .append('g')
      .attr('class', yClass)
      .attr(
        'transform',
        `translate(
          ${this.margin.left},
          ${this.margin.top}
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
    xTicks = 5,
    yTicks = 5,
    xTickFormat = () => '',
    yTickFormat = () => '',
  }: D3AxisChartSetAxisBackgroundGridParams) {
    this.axisGridXTicks = xTicks;
    this.axisGridYTicks = yTicks;

    if (direction.x) {
      this.axisGridXSvg = this.svg
        .append('g')
        .attr('class', xClass)
        .attr(
          'transform',
          `translate(
            ${this.margin.left},
            ${this.margin.top}
          )`,
        )
        .style('font-size', this.axisFontSize);

      this.axisGridX = axisRight(this.yScale())
        .tickSize(this.width)
        .ticks(this.axisGridXTicks)
        .tickFormat(xTickFormat);
    }

    if (direction.y) {
      this.axisGridYSvg = this.svg
        .append('g')
        .attr('class', yClass)
        .attr(
          'transform',
          `translate(
            ${this.margin.left},
            ${this.height - this.margin.top}
        )`,
        )
        .style('fornt-size', this.axisFontSize);

      this.axisGridY = axisTop(this.xScale())
        .tickSize(this.height)
        .ticks(this.axisGridYTicks)
        .tickFormat(yTickFormat);
    }
  }

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

  drawCircle() {}

  drawLine({
    data,
    color = 'black',
    strokeWidth = 2,
    lineType = 'STRAIGHT',
    animate = false,
    duration = 1500,
  }: D3AxisChartDrawLineParams) {
    this.strokeWidth = strokeWidth;
    this.lineType = lineType;

    const xScale = this.xScale();
    const yScale = this.yScale();

    const lineGenerator = line()
      .x((d) => xScale(d[0]))
      .y((d) => yScale(d[1]));

    if (this.lineType === 'CURVE') {
      lineGenerator.curve(curveBasis);
    }

    const path = this.svg
      .append('path')
      .attr('fill', 'none')
      .attr('d', `${lineGenerator(data.d3Position)}`)
      .attr('stroke-width', this.strokeWidth)
      .attr('stroke', color)
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.strokeWidth},
          ${this.margin.top}
        )
        `,
      )
      .attr('class', `${this.lineKey}${this.replaceColorHex(color)}`)
      .on('mouseover', (d) => {
        const targetClass = d.target.classList[0];
        const colorKey = targetClass.split('-')[1];

        select(`.${this.lineKey}${colorKey}`)
          .style('stroke-width', this.strokeWidth * 2.5)
          .style('cursor', 'pointer');

        select(`.${this.areaKey}${colorKey}`)
          .style('fill-opacity', 0.7)
          .style('cursor', 'pointer');
      })
      .on('mouseout', (d) => {
        const targetClass = d.target.classList[0];
        const colorKey = targetClass.split('-')[1];

        select(`.${this.lineKey}${colorKey}`)
          .style('stroke-width', this.strokeWidth)
          .style('cursor', 'default');

        select(`.${this.areaKey}${colorKey}`)
          .style('fill-opacity', 0)
          .style('cursor', 'pointer');
      });

    const pathLength = path.node()?.getTotalLength();

    if (animate && path && pathLength) {
      path
        .attr('stroke-dashoffset', pathLength)
        .attr('stroke-dasharray', pathLength)
        .transition()
        .ease(easeSinInOut)
        .duration(duration)
        .attr('stroke-dashoffset', 0); //시작점
    }
  }

  drawArea({
    data,
    color = 'black',
    opacity = 0,
    animate = false,
    duration = 1500,
  }: D3AxisChartDrawAreaParams) {
    const xScale = this.xScale();
    const yScale = this.yScale();

    const areaGenerator = area()
      .x((d) => xScale(d[0]))
      .y0(this.yRange[0])
      .y1((d) => yScale(d[1]));

    if (this.lineType === 'CURVE') {
      areaGenerator.curve(curveBasis);
    }

    const path = this.svg
      .append('path')
      .attr('fill', color)
      .attr('fill-opacity', opacity)
      .attr('stroke', 'none')
      .attr('d', `${areaGenerator(data.d3Position)}`)
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.strokeWidth},
          ${this.margin.top}
        )
        `,
      )
      .attr('class', `${this.areaKey}${this.replaceColorHex(color)}`)
      .on('mouseover', (d) => {
        const targetClass = d.target.classList[0];
        const colorKey = targetClass.split('-')[1];

        select(`.${this.lineKey}${colorKey}`)
          .style('stroke-width', this.strokeWidth * 2.5)
          .style('cursor', 'pointer');

        select(`.${this.areaKey}${colorKey}`)
          .style('fill-opacity', 0.7)
          .style('cursor', 'pointer');
      })
      .on('mouseout', (d) => {
        const targetClass = d.target.classList[0];
        const colorKey = targetClass.split('-')[1];

        select(`.${this.lineKey}${colorKey}`)
          .style('stroke-width', this.strokeWidth)
          .style('cursor', 'default');

        select(`.${this.areaKey}${colorKey}`)
          .style('fill-opacity', 0)
          .style('cursor', 'pointer');
      });

    const pathLength = path.node()?.getTotalLength();

    if (animate && path && pathLength) {
      path
        .attr('fill-opacity', 0)
        .transition()
        .ease(easeSinInOut)
        .duration(duration)
        .attr('fill-opacity', opacity);
    }
  }
}
