import {
  area,
  axisBottom,
  axisLeft,
  curveBasis,
  curveMonotoneX,
  curveMonotoneY,
  easeSinInOut,
  extent,
  line,
  scaleLinear,
  scaleTime,
} from 'd3';
import { startOfMonth, endOfMonth } from 'date-fns';
import { v4 as uuidv4 } from 'uuid';
import { palette } from '../../../../styles';
import { getRandomColors } from '../../../utils';
import D3Common2 from '../D3Common2';

import {
  D3Axis,
  D3Data,
  D3Domain,
  D3DoubleNumberArray,
  D3Margin,
  D3ScaleType,
  D3SelectionSVG,
  D3TickFormat,
} from '../D3Common2/D3Common2Types';
import { D3ValidationUtil } from '../D3Util';

import {
  D3AxisChartAreaType,
  D3AxisChartConstructorParams,
  D3AxisChartLinecapType,
  D3AxisChartLinecurvKeys,
  D3AxisChartLinecurvType,
  D3AxisChartLinejoinType,
  D3AxisChartLineType,
  D3AxisChartSetAreaOptionsParams,
  D3AxisChartSetAxisOptionsParams,
  D3AxisChartSetCircleOptionsParams,
  D3AxisChartSetLineOptionsParams,
} from './D3AxisChart2Types';

export default class D3AxisChart2 extends D3Common2 {
  /**
   * constructor
   */
  private svg: D3SelectionSVG;
  private width: number;
  private height: number;
  private margin: D3Margin;
  private xRange: D3DoubleNumberArray;
  private yRange: D3DoubleNumberArray;

  /**
   * data
   */
  private data: D3Data[][] = [];
  private xDomainKey = 'x';
  private yDomainKey = 'y';
  private xDomain: D3Domain = [0, 0];
  private yDomain: D3Domain = [0, 0];

  /**
   * scale
   */
  private xScaleType: D3ScaleType = 'number';
  private yScaleType: D3ScaleType = 'number';

  /**
   * axis options
   */
  private axisXTicks = 0;
  private axisYTicks = 0;
  private axisXTickSize = 0;
  private axisYTickSize = 0;
  private axisXTickFormat: D3TickFormat = (d, i) => `${d}`;
  private axisYTickFormat: D3TickFormat = (d, i) => `${d}`;
  private axisXClass = '';
  private axisYClass = '';
  private axisTransitionDuration = 850;

  /**
   * axis
   */
  private axisX: D3Axis | null = null;
  private axisY: D3Axis | null = null;

  /**
   * uniq class and color values
   */
  private uniqIdentifierValueMap: Map<number, Record<'key' | 'color', string>> =
    new Map();

  /**
   * line options
   */
  private readonly linecurvTypeMap: Map<
    D3AxisChartLinecurvKeys,
    typeof curveBasis | typeof curveMonotoneX | typeof curveMonotoneY
  > = new Map([
    ['curveBasis', curveBasis],
    ['curveMonotoneX', curveMonotoneX],
    ['curveMonotoneY', curveMonotoneY],
  ]);

  private lineType: D3AxisChartLineType = 'STRAIGHT';
  private linecurvType: D3AxisChartLinecurvType = curveMonotoneX;
  private linecapType: D3AxisChartLinecapType = 'butt';
  private linejoinType: D3AxisChartLinejoinType = 'miter';
  private lineStrokeWidth = 1;
  private lineTransition = true;
  private lineTransitionDuration = 550;

  /**
   * area options
   */
  private areaType: D3AxisChartAreaType = 'full';
  private areaOpacity = 0.2;
  private areaTransitionDuration = 550;

  /**
   * circle options
   */
  private circleRadius: number = 3;
  private circleStrokeWidth: number = 2;
  private circleIsFill: boolean = false;
  private circleTransitionDuration = 550;

  constructor({
    container,
    width,
    height,
    className = '',
    margin,
  }: D3AxisChartConstructorParams) {
    super();
    console.log('event: initialize');

    this.svg = this.appendSVG({
      container,
      width,
      height,
      className,
    });
    this.width = width;
    this.height = height;
    this.margin = margin;
    this.xRange = [0, width - (margin.left + margin.right)];
    this.yRange = [height - (margin.top + margin.bottom), 0];
  }

  setUniqIdentifierValueMap() {
    console.log('event: setUniqIdentifierValueMap');

    const colors = getRandomColors(this.data.length);
    this.data.forEach((_, i) => {
      const key = uuidv4();
      const color = colors[i];
      this.uniqIdentifierValueMap.set(i, { key, color });
    });
  }

  clearUniqIdentifierValueMap() {
    console.log('event: clearUniqIdentifierValueMap');

    this.uniqIdentifierValueMap.clear();
  }

  setData(data: D3Data[][]) {
    console.log('event: setData');

    this.data = data;
  }

  resetData() {
    console.log('event: resetData');

    this.data = [];
  }

  setScaleType(xType: D3ScaleType, yType: D3ScaleType) {
    console.log('event: setScaleType');

    this.xScaleType = xType;
    this.yScaleType = yType;
  }

  setDomainOptions(xDomainKey?: string, yDomainKey?: string) {
    console.log('event: setDomainOptions');

    if (xDomainKey) this.xDomainKey = xDomainKey;
    if (yDomainKey) this.yDomainKey = yDomainKey;
  }

  // filter data and calculate min, max domain data
  private getDomainMinMax() {
    console.log('event: getDomainMinMax');

    /**
     * flatten the data like this
     * [
     *   [{...}, {...}, {...}]
     *   [{...}, {...}, {...}]
     * ]
     */
    const flatData = this.data.flat();

    const xDomainPool = [];
    const yDomainPool = [];

    // there may be a lot of data, so write it with a for loop
    for (let i = 0; i < flatData.length; i++) {
      xDomainPool.push(flatData[i][this.xDomainKey]);
      yDomainPool.push(flatData[i][this.yDomainKey]);
    }

    const [xMin = 0, xMax = 0] = extent(xDomainPool);
    const [yMin = 0, yMax = 0] = extent(yDomainPool);

    return {
      xMin,
      xMax,
      yMin,
      yMax,
    };
  }

  private setDomainData({
    xMin,
    xMax,
    yMin,
    yMax,
  }: ReturnType<typeof this.getDomainMinMax>) {
    console.log('event: setDomainData');

    // xScale
    if (this.xScaleType === 'number') {
      const max = this.calcMaxOfNumber(xMax);
      this.xDomain = [0, max];
    }
    if (this.xScaleType === 'time') {
      const min = startOfMonth(xMin);
      const max = endOfMonth(xMax);
      this.xDomain = [min, max];
    }

    // yScale
    if (this.yScaleType === 'number') {
      const max = this.calcMaxOfNumber(yMax);
      this.yDomain = [0, max];
    }
    if (this.yScaleType === 'time') {
      const min = startOfMonth(yMin);
      const max = endOfMonth(yMax);
      this.yDomain = [min, max];
    }
  }

  // set domain data
  setDomain() {
    console.log('event: setDomain');
    this.setDomainData({ ...this.getDomainMinMax() });
  }

  private xScale() {
    if (this.xScaleType === 'time') {
      return scaleTime().domain(this.xDomain).range(this.xRange).nice();
    }
    // axis type number;
    return scaleLinear().domain(this.xDomain).range(this.xRange).nice();
  }

  private yScale() {
    if (this.yScaleType === 'time') {
      return scaleTime().domain(this.yDomain).range(this.yRange).nice();
    }
    // axis type number;
    return scaleLinear().domain(this.yDomain).range(this.yRange).nice();
  }

  setAxisOptions({
    axisXTicks,
    axisYTicks,
    axisXTickVisible,
    axisYTickVisible,
    axisXTickFormat,
    axisYTickFormat,
    axisXClass,
    axisYClass,
    axisTransitionDuration,
  }: D3AxisChartSetAxisOptionsParams) {
    console.log('event: setAxisOptions');

    if (axisXTicks) this.axisXTicks = axisXTicks;
    if (axisYTicks) this.axisYTicks = axisYTicks;
    if (axisXTickFormat) this.axisXTickFormat = axisXTickFormat;
    if (axisYTickFormat) this.axisYTickFormat = axisYTickFormat;
    if (axisXClass) this.axisXClass = axisXClass;
    if (axisYClass) this.axisYClass = axisYClass;
    if (axisTransitionDuration)
      this.axisTransitionDuration = axisTransitionDuration;

    if (axisXTickVisible && this.height) {
      this.axisXTickSize = this.height - (this.margin.bottom + this.margin.top);
    }
    if (axisYTickVisible && this.width) {
      this.axisYTickSize =
        this.width - (this.margin.left + this.margin.right * 0.4);
    }
  }

  setAxis() {
    console.log('event: setAxis');

    this.axisX = axisBottom(this.xScale())
      .tickSize(0)
      .tickSizeInner(-this.axisXTickSize)
      .ticks(this.axisXTicks)
      .tickFormat(this.axisXTickFormat);

    this.axisY = axisLeft(this.yScale())
      .tickSize(0)
      .tickSizeInner(this.axisYTickSize)
      .ticks(this.axisYTicks)
      .tickFormat(this.axisYTickFormat);
  }

  appendAxis() {
    console.log('event: appendAxis');

    if (this.axisX) {
      this.svg
        .append('g')
        .attr('class', this.axisXClass)
        .attr(
          'transform',
          `translate(
            ${this.margin.left + this.margin.right * 0.4}, 
            ${this.height - this.margin.top}
          )`,
        )
        .call(this.axisX);
    }

    if (this.axisY) {
      this.svg
        .append('g')
        .attr('class', this.axisYClass)
        .attr(
          'transform',
          `translate(
            ${this.width - (this.margin.left - this.margin.right * 0.4)}
            ${this.margin.top}
          )`,
        )
        .call(this.axisY);
    }
  }

  /**
   * previously possible actions
   * @this.setData
   * @this.setScaleType
   * @this.setDomainOptions
   * @this.setDomain
   * @this.setAxisOptions
   * @this.setAxis
   */
  updateAxis() {
    console.log('event: updateAxis');

    if (this.axisX) {
      this.svg
        .selectAll(`.${this.axisXClass}`)
        .transition()
        .duration(this.axisTransitionDuration)
        .call(this.axisX as any);
    }

    if (this.axisY) {
      this.svg
        .selectAll(`.${this.axisYClass}`)
        .transition()
        .duration(this.axisTransitionDuration)
        .call(this.axisY as any);
    }
  }

  setLinecurvtype(key: D3AxisChartLinecurvKeys) {
    const curvtype = this.linecurvTypeMap.get(key) ?? curveMonotoneX;
    this.linecurvType = curvtype;
  }

  setLineOptions({
    lineType,
    linecurvType,
    linecapType,
    linejoinType,
    lineStrokeWidth,
    lineTransition,
    lineTransitionDuration,
  }: D3AxisChartSetLineOptionsParams) {
    console.log('event: setLineOptions');

    if (lineType) this.lineType = lineType;
    if (linecurvType) this.setLinecurvtype(linecurvType);
    if (linecapType) this.linecapType = linecapType;
    if (linejoinType) this.linejoinType = linejoinType;
    if (lineStrokeWidth) this.lineStrokeWidth = lineStrokeWidth;
    if (lineTransition) this.lineTransition = lineTransition;
    if (lineTransitionDuration)
      this.lineTransitionDuration = lineTransitionDuration;
  }

  private getLineGenerator() {
    console.log('event: getLineGenerator');

    const lineGenerator = line<D3Data>()
      .x((d) => this.xScale()(d[this.xDomainKey]))
      .y((d) => this.yScale()(d[this.yDomainKey]));

    if (this.lineType === 'CURVE') {
      lineGenerator.curve(this.linecurvType);
    }
    return lineGenerator;
  }

  appendLine() {
    console.log('event: appendLine');

    D3ValidationUtil.isExistMapValidate(
      this.uniqIdentifierValueMap,
      'the setUniqIdentifierValueMap function must be executed first before the action to draw the chart.',
    );

    const lineGenerator = this.getLineGenerator();

    this.data.forEach((data, i) => {
      const map = this.uniqIdentifierValueMap.get(i);
      const key = map!.key;
      const color = map!.color;
      this.svg
        .append('path')
        .attr('fill', 'none')
        .attr('stroke-width', this.lineStrokeWidth)
        .attr('stroke', color)
        .attr('stroke-linejoin', this.linejoinType)
        .attr('stroke-linecap', this.linecapType)
        .attr('class', `line-${key}`)
        .attr(
          'transform',
          `translate(
            ${this.margin.left + this.margin.right * 0.4},
            ${this.margin.top}
          )`,
        )
        .attr('d', `${lineGenerator(data)}`);

      // const pathLength = path.node()?.getTotalLength();
      // if (this.lineTransition) {
      //   path
      //     .attr('stroke-dashoffset', pathLength)
      //     .attr('stroke-dasharray', pathLength)
      //     .transition()
      //     .ease(easeSinInOut)
      //     .duration(this.lineTransitionDuration)
      //     .attr('stroke-dashoffset', 0);
      // }
    });
  }

  updateLine() {
    console.log('event: updateLine');

    const lineGenerator = this.getLineGenerator();

    this.data.forEach((data, i) => {
      const map = this.uniqIdentifierValueMap.get(i);
      const key = map!.key;
      this.svg
        .selectAll(`.line-${key}`)
        .transition()
        .ease(easeSinInOut)
        .duration(this.lineTransitionDuration)
        .attr('d', `${lineGenerator(data)}`);
    });
  }

  setAreaOptions({
    areaType,
    areaOpacity,
    areaTransitionDuration,
  }: D3AxisChartSetAreaOptionsParams) {
    console.log('event: setAreaOptions');

    if (areaType) this.areaType = areaType;
    if (areaOpacity) this.areaOpacity = areaOpacity;
    if (areaTransitionDuration)
      this.areaTransitionDuration = this.areaTransitionDuration;
  }

  private getAreaGenerator() {
    console.log('event: getAreaGenerator');

    const areaGenerator = area<D3Data>();

    // area
    if (this.areaType === 'full') {
      areaGenerator
        .x0((d) => this.xScale()(d[this.xDomainKey]))
        .y0(this.yRange[0])
        .y1((d) => this.yScale()(d[this.yDomainKey]));
    }
    if (this.areaType === 'boundary') {
      areaGenerator
        .x1((d) => this.xScale()(d[this.xDomainKey]))
        .y0(this.yRange[0])
        .y1((d) => this.yScale()(d[this.yDomainKey]));
    }

    // curv
    if (this.lineType === 'CURVE') {
      areaGenerator.curve(this.linecurvType);
    }

    return areaGenerator;
  }

  appendArea() {
    console.log('event: appendArea');

    D3ValidationUtil.isExistMapValidate(
      this.uniqIdentifierValueMap,
      'the setUniqIdentifierValueMap function must be executed first before the action to draw the chart.',
    );

    const areaGenerator = this.getAreaGenerator();

    this.data.forEach((data, i) => {
      const map = this.uniqIdentifierValueMap.get(i);
      const key = map!.key;
      const color = map!.color;
      this.svg
        .append('path')
        .attr('fill', color)
        .attr('fill-opacity', this.areaOpacity)
        .attr('stroke', 'none')
        .attr('class', `area-${key}`)
        .attr(
          'transform',
          `translate(
            ${this.margin.left + this.margin.right * 0.4},
            ${this.margin.top}
          )`,
        )
        .attr('d', `${areaGenerator(data)}`);
    });
  }

  updateArea() {
    const areaGenerator = this.getAreaGenerator();

    this.data.forEach((data, i) => {
      const map = this.uniqIdentifierValueMap.get(i);
      const key = map!.key;
      this.svg
        .selectAll(`.area-${key}`)
        .transition()
        .ease(easeSinInOut)
        .duration(this.lineTransitionDuration)
        .attr('d', `${areaGenerator(data)}`);
    });
  }

  setCircleOptions({
    circleRadius,
    circleStrokeWidth,
    circleIsFill,
  }: D3AxisChartSetCircleOptionsParams) {
    console.log('event: setCircleOptions');

    if (circleRadius) this.circleRadius = circleRadius;
    if (circleStrokeWidth) this.circleStrokeWidth = circleStrokeWidth;
    if (circleIsFill) this.circleIsFill = circleIsFill;
  }

  appendCircle() {
    console.log('event: appendArea');

    D3ValidationUtil.isExistMapValidate(
      this.uniqIdentifierValueMap,
      'the setUniqIdentifierValueMap function must be executed first before the action to draw the chart.',
    );

    this.svg.selectAll('.circle').remove();

    this.data.forEach((data, i) => {
      const map = this.uniqIdentifierValueMap.get(i);
      const color = map!.color;

      this.svg
        .selectAll('circles')
        .data(data)
        .join('circle')
        .attr('fill', this.circleIsFill ? color : palette.white)
        .attr('stroke', color)
        .attr('stroke-width', this.circleStrokeWidth)
        .attr('r', this.circleRadius)
        .attr('class', `circle`)
        .attr(
          'transform',
          `translate(
            ${this.margin.left + this.margin.right * 0.4},
            ${this.margin.top}
          )`,
        )
        .attr('cx', (d: any) => this.xScale()(d[this.xDomainKey]))
        .attr('cy', (d: any) => this.yScale()(d[this.yDomainKey]));
    });
  }
}
