import { scaleLinear } from 'd3';
import D3Common, { D3CommonInterface } from './D3Common';
import { D3DoubleNumberArray, D3Selection } from './types/d3CommonTypes';
import { D3LineChartConstructorParams } from './types/d3LineChartTypes';

export default class D3LineChart extends D3Common implements D3CommonInterface {
  private svg: D3Selection;
  private xDomain: D3DoubleNumberArray;
  private yDomain: D3DoubleNumberArray;
  private xRange: D3DoubleNumberArray;
  private yRange: D3DoubleNumberArray;

  constructor({
    container,
    width,
    height,
    className = '',
    xDomainData,
    yDomainData,
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
    this.xDomain = this.getExtent(xDomainData);
    this.yDomain = this.getExtent(yDomainData);
    this.xRange = xRange;
    this.yRange = yRange;
  }

  private xScale() {
    return scaleLinear().domain(this.xDomain).range(this.xRange);
  }

  private yScale() {
    return scaleLinear().domain(this.yDomain).range(this.yRange);
  }

  setAxis() {}
}
