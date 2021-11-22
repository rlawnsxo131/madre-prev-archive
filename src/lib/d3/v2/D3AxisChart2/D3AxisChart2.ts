import { extent, scaleLinear, scaleTime } from 'd3';
import D3Common2 from '../D3Common2';
import {
  D3Data,
  D3DoubleNumberArray,
  D3Margin,
  D3ScaleType,
  D3Selection,
  D3SelectionSVG,
} from '../D3Common2/D3Common2Types';
import {
  D3AxisChartConstructorParams,
  D3AxisChartSetAxisOptionsParams,
  D3AxisChartSetDataAndDomainParams,
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
  private data: D3Data[] = [];
  private xDomain: D3DoubleNumberArray = [0, 0];
  private yDomain: D3DoubleNumberArray = [0, 0];

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
  private axisXClass = '';
  private axisYClass = '';
  private axisFontSize = 10;

  /**
   * axis svg
   */
  private axisXSVG: D3Selection | null = null;
  private axisYSVG: D3Selection | null = null;

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
    this.xDomain = [xMin, xMax];
    this.yDomain = [yMin, yMax];
  }

  setDataAndDomain({ data, xKey, yKey }: D3AxisChartSetDataAndDomainParams) {
    this.setData(data);
    this.setDomain(xKey, yKey);
  }

  setScaleType(xType: D3ScaleType, yType: D3ScaleType) {
    this.xScaleType = xType;
    this.yScaleType = yType;
  }

  private xScale() {
    if (this.xScaleType === 'number') {
      return scaleLinear().domain(this.xDomain).range(this.xRange).nice();
    }
    if (this.xScaleType === 'time') {
      return scaleTime().domain(this.xDomain).range(this.xRange).nice();
    }
  }

  private yScale() {
    if (this.yScaleType === 'number') {
      return scaleLinear().domain(this.yDomain).range(this.yRange).nice();
    }
    if (this.yScaleType === 'time') {
      return scaleTime().domain(this.yDomain).range(this.yRange).nice();
    }
  }

  setAxisOptions({
    axisXTicks = 0,
    axisYTicks = 0,
    axisXTickSize = 0,
    axisYTickSize = 0,
    axisXClass = '',
    axisYClass = '',
    axisFontSize = 10,
  }: D3AxisChartSetAxisOptionsParams) {
    this.axisXTicks = axisXTicks;
    this.axisYTicks = axisYTicks;
    this.axisXTickSize = axisXTickSize;
    this.axisYTickSize = axisYTickSize;
    this.axisXClass = axisXClass;
    this.axisYClass = axisYClass;
    this.axisFontSize = axisFontSize;
  }

  setAxisSvg() {
    this.axisXSVG = null;
    this.axisYSVG = null;
  }
}
