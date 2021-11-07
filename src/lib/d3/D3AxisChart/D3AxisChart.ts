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
  D3AxisChartDrawCircleParams,
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
  private axisFontSize: number = 10;
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
   * @AxisLine @AxisArea
   */
  private readonly lineKey = 'line';
  private readonly areaKey = 'area';
  private readonly lineAndAreaKeyRegex = /(line-|area-)/gi;
  private lineType: D3AxisChartLineType = 'STRAIGHT';
  private commonKeyMap: Map<string, string> = new Map([]);
  private mouseOverActionMap = new Map([
    [this.lineKey, false],
    [this.areaKey, false],
  ]);
  private mouseOverOpacity: number = 0.6;
  private strokeWidth: number = 2;

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

  private replaceLineOrAreaKey(target: string) {
    return target.replace(this.lineAndAreaKeyRegex, '');
  }

  private getCommonKeyAndClassName(
    type: typeof this.lineKey | typeof this.areaKey,
    color: string,
    uuid: string,
  ) {
    const replaceHexColor = this.replaceSharpFromHexColor(color);
    const commonKey = `${replaceHexColor}-${uuid}`;
    const className = `${type}-${commonKey}`; // key + color + uuid
    return { commonKey, className };
  }

  private setCommonKeyMap(commonKey: string, color: string) {
    const defined = !!this.commonKeyMap.get(commonKey);
    if (!defined) {
      this.commonKeyMap.set(commonKey, color);
    }
  }

  private onMouseOverAction(targetClass: string) {
    if (this.mouseOverActionMap.get(this.lineKey)) {
      select(`.${this.lineKey}-${targetClass}`)
        .style('stroke-width', this.strokeWidth * 2)
        .style('cursor', 'pointer');
    }
    if (this.mouseOverActionMap.get(this.areaKey)) {
      select(`.${this.areaKey}-${targetClass}`)
        .style('fill-opacity', this.mouseOverOpacity)
        .style('cursor', 'pointer');
    }
  }

  private onMouseOutAction(targetClass: string) {
    if (this.mouseOverActionMap.get(this.lineKey)) {
      select(`.${this.lineKey}-${targetClass}`)
        .style('stroke-width', this.strokeWidth)
        .style('cursor', 'default');
    }
    if (this.mouseOverActionMap.get(this.areaKey)) {
      select(`.${this.areaKey}-${targetClass}`)
        .style('fill-opacity', 0)
        .style('cursor', 'pointer');
    }
  }

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

    this.axisXSvg = this.svg
      .append('g')
      .attr('class', xClass)
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.margin.left * 0.4}, 
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
      .tickSize(xTickSize)
      .ticks(xTicks)
      .tickFormat(xTickFormat);

    this.axisY = axisLeft(this.yScale())
      .tickSize(yTickSize)
      .ticks(yTicks)
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
        );

      this.axisGridX = axisRight(this.yScale())
        .tickSize(this.width)
        .ticks(xTicks)
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
        );

      this.axisGridY = axisTop(this.xScale())
        .tickSize(this.height)
        .ticks(yTicks)
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

  drawCircle({
    data,
    color = 'black',
    radius = 3,
    uuid = '',
    isMouseOverAction = false,
  }: D3AxisChartDrawCircleParams) {
    const xScale = this.xScale();
    const yScale = this.yScale();
    this.svg
      .selectAll('circles')
      .data(data.d3Position)
      .enter()
      .append('circle')
      .attr('fill', color)
      .attr('r', radius)
      .attr('cx', (d) => xScale(d[0]))
      .attr('cy', (d) => yScale(d[1]))
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.margin.left * 0.4},
          ${this.margin.top}
        )
        `,
      );
  }

  drawLine({
    data,
    color = 'black',
    strokeWidth = 2,
    lineType = 'STRAIGHT',
    animate = false,
    duration = 1500,
    uuid = '',
    linejoinType = 'miter',
    linecapType = 'butt',
    isMouseOverAction = false,
  }: D3AxisChartDrawLineParams) {
    const { commonKey, className } = this.getCommonKeyAndClassName(
      this.lineKey,
      color,
      uuid,
    );
    this.setCommonKeyMap(commonKey, color);
    this.mouseOverActionMap.set(this.lineKey, isMouseOverAction);

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
      .attr('stroke-linejoin', linejoinType)
      .attr('stroke-linecap', linecapType)
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.margin.left * 0.4},
          ${this.margin.top}
        )
        `,
      )
      .attr('class', className);

    if (this.mouseOverActionMap.get(this.lineKey)) {
      path.on('mouseover', (d) => {
        const targetClass = this.replaceLineOrAreaKey(d.target.classList[0]);
        this.onMouseOverAction(targetClass);
      });
      path.on('mouseout', (d) => {
        const targetClass = this.replaceLineOrAreaKey(d.target.classList[0]);
        this.onMouseOutAction(targetClass);
      });
    }

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
    uuid = '',
    areaType = 'full',
    isMouseOverAction = false,
    mouseOverOpacity = 0.6,
  }: D3AxisChartDrawAreaParams) {
    const { commonKey, className } = this.getCommonKeyAndClassName(
      this.areaKey,
      color,
      uuid,
    );
    this.mouseOverOpacity = mouseOverOpacity;
    this.setCommonKeyMap(commonKey, color);
    this.mouseOverActionMap.set(this.areaKey, isMouseOverAction);

    const xScale = this.xScale();
    const yScale = this.yScale();

    const areaGenerator = area();

    if (areaType === 'full') {
      areaGenerator
        .x0((d) => xScale(d[0]))
        .y0(this.yRange[0])
        .y1((d) => yScale(d[1]));
    }
    if (areaType === 'boundary') {
      areaGenerator
        .x1((d) => xScale(d[0]))
        .y0(this.yRange[0])
        .y1((d) => yScale(d[1]));
    }
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
          ${this.margin.left + this.margin.left * 0.4},
          ${this.margin.top}
        )
        `,
      )
      .attr('class', className);

    if (isMouseOverAction) {
      path.on('mouseover', (d) => {
        const targetClass = this.replaceLineOrAreaKey(d.target.classList[0]);
        this.onMouseOverAction(targetClass);
      });
      path.on('mouseout', (d) => {
        const targetClass = this.replaceLineOrAreaKey(d.target.classList[0]);
        this.onMouseOutAction(targetClass);
      });
    }

    // const pathLength = path.node()?.getTotalLength();

    if (animate && path) {
      path
        .attr('fill-opacity', 0)
        .transition()
        .ease(easeSinInOut)
        .duration(duration)
        .attr('fill-opacity', opacity);
    }
  }
}
