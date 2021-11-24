import { axisBottom, axisLeft, extent, scaleLinear, scaleTime } from 'd3';
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
import {
  D3AxisChartConstructorParams,
  D3AxisChartSetAxisOptionsParams,
} from './D3AxisChart2Types';
import { startOfMonth, endOfMonth } from 'date-fns';
import { palette } from '../../../../styles';

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
  private data: D3Data[] = [];
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
  private axisFontSize = 10;

  /**
   * axis svg
   */
  private axisX: D3Axis | null = null;
  private axisY: D3Axis | null = null;

  constructor({
    container,
    width,
    height,
    className = '',
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
    this.margin = margin;
    this.xRange = [0, width - (margin.left + margin.right)];
    this.yRange = [height - (margin.top + margin.bottom), 0];
  }

  setData(data: D3Data[]) {
    this.data = data;
  }

  setScaleType(xType: D3ScaleType, yType: D3ScaleType) {
    this.xScaleType = xType;
    this.yScaleType = yType;
  }

  /**
   * There may be a lot of data, so write it with a for loop
   */
  setDomain(xKey: string, yKey: string) {
    const xDomainPool = [];
    const yDomainPool = [];
    for (let i = 0; i < this.data.length; i++) {
      xDomainPool.push(this.data[i][xKey]);
      yDomainPool.push(this.data[i][yKey]);
    }
    const [xMin = 0, xMax = 0] = extent(xDomainPool);
    const [yMin = 0, yMax = 0] = extent(yDomainPool);

    if (this.xScaleType === 'number') {
      const max = Math.ceil(xMax / 100) * 100;
      this.xDomain = [0, max];
    } else {
      const min = startOfMonth(xMin);
      const max = endOfMonth(xMax);
      this.xDomain = [min, max];
    }

    if (this.yScaleType === 'number') {
      const max = Math.ceil(yMax / 100) * 100;
      this.yDomain = [0, max];
    } else {
      const min = startOfMonth(yMin);
      const max = endOfMonth(yMax);
      this.yDomain = [min, max];
    }
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
    axisXTicks = 0,
    axisYTicks = 0,
    axisXTickSize = 0,
    axisYTickSize = 0,
    axisXTickFormat = (d, i) => `${d}`,
    axisYTickFormat = (d, i) => `${d}`,
    axisXClass = '',
    axisYClass = '',
    axisFontSize = 10,
  }: D3AxisChartSetAxisOptionsParams) {
    this.axisXTicks = axisXTicks;
    this.axisYTicks = axisYTicks;
    this.axisXTickSize = axisXTickSize - (this.margin.bottom + this.margin.top);
    this.axisYTickSize =
      axisYTickSize - (this.margin.left + this.margin.right * 0.4);
    this.axisXTickFormat = axisXTickFormat;
    this.axisYTickFormat = axisYTickFormat;
    this.axisXClass = axisXClass;
    this.axisYClass = axisYClass;
    this.axisFontSize = axisFontSize;
  }

  setAxis() {
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
        .style('font-size', this.axisFontSize)
        .call(this.axisX)
        .selectAll('line, path')
        .style('color', palette.gray['300']);
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
        .style('font-size', this.axisFontSize)
        .call(this.axisY)
        .selectAll('line, path')
        .style('color', palette.gray['300']);
    }
  }
}
