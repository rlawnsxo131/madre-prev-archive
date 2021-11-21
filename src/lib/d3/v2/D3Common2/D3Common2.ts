import { select } from 'd3';
import { AppendSVGParams } from './D3Common2Types';

export default class D3Common {
  protected appendSvg({
    container,
    width,
    height,
    className = '',
  }: AppendSVGParams) {
    const svg = select(container)
      .append('svg')
      .attr('width', width)
      .attr('height', height)
      .attr('viewBox', `0, 0, ${width}, ${height}`)
      .attr('class', className);

    return svg;
  }
}
