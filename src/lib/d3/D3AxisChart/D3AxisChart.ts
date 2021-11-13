import {
  area,
  axisBottom,
  axisLeft,
  axisRight,
  axisTop,
  curveBasis,
  curveMonotoneX,
  curveMonotoneY,
  easeSinInOut,
  extent,
  line,
  scaleLinear,
  select,
} from 'd3';
import { palette } from '../../../styles';
import D3Common from '../D3Common';
import {
  D3Axis,
  D3DoubleNumberArray,
  D3Selection,
  D3SelectionSVG,
} from '../D3Common/D3CommonTypes';
import {
  D3AxisChartConstructorParams,
  D3AxisChartCurvType,
  D3AxisChartDrawAreaParams,
  D3AxisChartDrawAxisParams,
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
  private axisXTicks: number = 0;
  private axisYTicks: number = 0;
  private axisXTickSize: number = 0;
  private axisYTickSize: number = 0;
  private axisXSvg: D3Selection = null;
  private axisYSvg: D3Selection = null;

  /**
   * @AxisGrid
   */
  private axisGridXSvg: D3Selection = null;
  private axisGridYSvg: D3Selection = null;
  private axisGridX: D3Axis = null;
  private axisGridY: D3Axis = null;

  /**
   * @Line
   */
  private lineStrokeWidth: number = 2;

  /**
   * @Area
   */
  private areaMouseOverOpacity: number = 0.6;

  /**
   * @Circle
   */

  /**
   * @AxisLine @AxisArea @Circle
   */
  private readonly lineKey = 'line';
  private readonly areaKey = 'area';
  private readonly circleKey = 'circle';
  private readonly lineAndAreaKeyRegex = /(line-|area-|circle-)/gi;
  private lineType: D3AxisChartLineType = 'STRAIGHT';
  private lineCurvType: D3AxisChartCurvType = 'curveBasis';
  private readonly lineCurvTypeMap: Map<
    D3AxisChartCurvType,
    typeof curveBasis | typeof curveMonotoneX | typeof curveMonotoneY
  > = new Map([
    ['curveBasis', curveBasis],
    ['curveMonotoneX', curveMonotoneX],
    ['curveMonotoneY', curveMonotoneY],
  ]);
  private commonKeyMap: Map<string, string> = new Map([]);
  private mouseOverActionMap = new Map([
    [this.lineKey, false],
    [this.areaKey, false],
    [this.circleKey, false],
  ]);

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

  private getCurvType() {
    return (
      this.lineCurvTypeMap.get(this.lineCurvType) ??
      this.lineCurvTypeMap.get('curveBasis')!
    );
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
        .style('stroke-width', this.lineStrokeWidth * 2)
        .style('cursor', 'pointer');
    }
    if (this.mouseOverActionMap.get(this.areaKey)) {
      select(`.${this.areaKey}-${targetClass}`)
        .style('fill-opacity', this.areaMouseOverOpacity)
        .style('cursor', 'pointer');
    }
    if (this.mouseOverActionMap.get(this.circleKey)) {
    }
  }

  private onMouseOutAction(targetClass: string) {
    if (this.mouseOverActionMap.get(this.lineKey)) {
      select(`.${this.lineKey}-${targetClass}`)
        .style('stroke-width', this.lineStrokeWidth)
        .style('cursor', 'default');
    }
    if (this.mouseOverActionMap.get(this.areaKey)) {
      select(`.${this.areaKey}-${targetClass}`)
        .style('fill-opacity', 0)
        .style('cursor', 'pointer');
    }
    if (this.mouseOverActionMap.get(this.circleKey)) {
    }
  }

  setAxis({
    axisXTicks = 0,
    axisYTicks = 0,
    axisXTickSize = 0,
    axisYTickSize = 0,
    axisXClass = '',
    axisYClass = '',
    axisFontSize = 10,
  }: D3AxisChartSetAxisParams) {
    this.axisXTicks = axisXTicks;
    this.axisYTicks = axisYTicks;
    this.axisXTickSize = axisXTickSize;
    this.axisYTickSize = axisYTickSize;

    this.axisXSvg = this.svg
      .append('g')
      .attr('class', axisXClass)
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.margin.left * 0.4}, 
          ${this.height - this.margin.top}
        )`,
      )
      .style('font-size', axisFontSize);

    this.axisYSvg = this.svg
      .append('g')
      .attr('class', axisYClass)
      .attr(
        'transform',
        `translate(
          ${this.margin.left},
          ${this.margin.top}
        )`,
      )
      .style('font-size', axisFontSize);
  }

  drawAxis({
    xTickFormat = (d, i) => `${d}`,
    yTickFormat = (d, i) => `${d}`,
  }: D3AxisChartDrawAxisParams) {
    if (this.axisXSvg) {
      const axisX = axisBottom(this.xScale())
        .tickSize(this.axisXTickSize)
        .ticks(this.axisXTicks)
        .tickFormat(xTickFormat);
      axisX(this.axisXSvg);
    }
    if (this.axisYSvg) {
      const axisY = axisLeft(this.yScale())
        .tickSize(this.axisYTickSize)
        .ticks(this.axisYTicks)
        .tickFormat(yTickFormat);
      axisY(this.axisYSvg);
    }
  }

  setAxisBackgroundGrid({
    axisBackgroundGridDirection,
    axisBackgroundGridXTicks = 5,
    axisBackgroundGridYTicks = 5,
    axisBackgroundGridXTickFormat = () => '',
    axisBackgroundGridYTickFormat = () => '',
    axisBackgroundGridXClass = '',
    axisBackgroundGridYClass = '',
  }: D3AxisChartSetAxisBackgroundGridParams) {
    if (axisBackgroundGridDirection.x) {
      this.axisGridXSvg = this.svg
        .append('g')
        .attr('class', axisBackgroundGridXClass)
        .attr(
          'transform',
          `translate(
            ${this.margin.left},
            ${this.margin.top}
          )`,
        );

      this.axisGridX = axisRight(this.yScale())
        .tickSize(this.width)
        .ticks(axisBackgroundGridXTicks)
        .tickFormat(axisBackgroundGridXTickFormat);
    }

    if (axisBackgroundGridDirection.y) {
      this.axisGridYSvg = this.svg
        .append('g')
        .attr('class', axisBackgroundGridYClass)
        .attr(
          'transform',
          `translate(
            ${this.margin.left},
            ${this.height - this.margin.top}
          )`,
        );

      this.axisGridY = axisTop(this.xScale())
        .tickSize(this.height)
        .ticks(axisBackgroundGridYTicks)
        .tickFormat(axisBackgroundGridYTickFormat);
    }
  }

  /**
   * @AxisBackground
   */
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
    lineType = 'STRAIGHT',
    lineCurvType = 'curveBasis',
    lineStrokeWidth = 2,
    linejoinType = 'miter',
    linecapType = 'butt',
    lineDrawAnimate = false,
    lineDrawAnimateDuration = 1500,
    isMouseOverAction = false,
    uuid = '',
  }: D3AxisChartDrawLineParams) {
    const { commonKey, className } = this.getCommonKeyAndClassName(
      this.lineKey,
      color,
      uuid,
    );
    this.setCommonKeyMap(commonKey, color);
    this.mouseOverActionMap.set(this.lineKey, isMouseOverAction);
    this.lineStrokeWidth = lineStrokeWidth;
    this.lineType = lineType;
    this.lineCurvType = lineCurvType;

    const xScale = this.xScale();
    const yScale = this.yScale();

    const lineGenerator = line()
      .x((d) => xScale(d[0]))
      .y((d) => yScale(d[1]));

    if (this.lineType === 'CURVE') {
      lineGenerator.curve(this.getCurvType());
    }

    const path = this.svg
      .append('path')
      .attr('fill', 'none')
      .attr('d', `${lineGenerator(data.d3Position)}`)
      .attr('stroke-width', this.lineStrokeWidth)
      .attr('stroke', color)
      .attr('stroke-linejoin', linejoinType)
      .attr('stroke-linecap', linecapType)
      .attr('class', className)
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.margin.right * 0.4},
          ${this.margin.top}
        )
        `,
      );

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

    if (lineDrawAnimate && pathLength) {
      path
        .attr('stroke-dashoffset', pathLength)
        .attr('stroke-dasharray', pathLength)
        .transition()
        .ease(easeSinInOut)
        .duration(lineDrawAnimateDuration)
        .attr('stroke-dashoffset', 0); //시작점
    }
  }

  drawArea({
    data,
    color = 'black',
    areaOpacity = 0,
    areaDrawAnimate = false,
    areaDrawAnimateDuration = 1500,
    areaType = 'full',
    isMouseOverAction = false,
    areaMouseOverOpacity = 0.6,
    uuid = '',
  }: D3AxisChartDrawAreaParams) {
    const { commonKey, className } = this.getCommonKeyAndClassName(
      this.areaKey,
      color,
      uuid,
    );
    this.setCommonKeyMap(commonKey, color);
    this.mouseOverActionMap.set(this.areaKey, isMouseOverAction);
    this.areaMouseOverOpacity = areaMouseOverOpacity;

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
      areaGenerator.curve(this.getCurvType());
    }

    const path = this.svg
      .append('path')
      .attr('fill', color)
      .attr('fill-opacity', areaOpacity)
      .attr('stroke', 'none')
      .attr('d', `${areaGenerator(data.d3Position)}`)
      .attr('class', className)
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.margin.right * 0.4},
          ${this.margin.top}
        )
        `,
      );

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

    if (areaDrawAnimate && path) {
      path
        .attr('fill-opacity', 0)
        .transition()
        .ease(easeSinInOut)
        .duration(areaDrawAnimateDuration)
        .attr('fill-opacity', areaOpacity);
    }
  }

  /**
   * need to add animation, fill and stroke-width operations
   */
  drawCircle({
    data,
    color = 'black',
    circleRadius = 3,
    uuid = '',
    isMouseOverAction = false,
  }: D3AxisChartDrawCircleParams) {
    const { commonKey, className } = this.getCommonKeyAndClassName(
      this.lineKey,
      color,
      uuid,
    );
    this.setCommonKeyMap(commonKey, color);
    this.mouseOverActionMap.set(this.circleKey, isMouseOverAction);

    const xScale = this.xScale();
    const yScale = this.yScale();

    this.svg
      .selectAll('circles')
      .data(data.d3Position)
      .enter()
      .append('circle')
      .attr('fill', palette.white)
      .attr('stroke', color)
      .attr('stroke-width', 2)
      .attr('r', circleRadius)
      .attr('cx', (d) => xScale(d[0]))
      .attr('cy', (d) => yScale(d[1]))
      .attr('class', className)
      .attr(
        'transform',
        `translate(
          ${this.margin.left + this.margin.right * 0.4},
          ${this.margin.top}
        )
        `,
      );
  }
}
