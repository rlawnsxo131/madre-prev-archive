import { extent, randomInt, randomUniform, select } from 'd3';
import { AppendSvgParams, D3DoubleNumberArray } from './d3CommonTypes';

export default class D3Common {
  protected appendSvg({
    container,
    width,
    height,
    className = '',
  }: AppendSvgParams) {
    const svg = select(container)
      .append('svg')
      .attr('width', width)
      .attr('height', height)
      .attr('class', className);

    return svg;
  }

  protected getExtent(data: number[]): D3DoubleNumberArray {
    const [min = 0, max = 0] = extent(data);
    return [min, max];
  }

  protected getRandomInt(min: number, max: number) {
    return randomInt(min, max)();
  }

  protected getRendomUniform(min: number, max: number) {
    return randomUniform(min, max)();
  }
}
