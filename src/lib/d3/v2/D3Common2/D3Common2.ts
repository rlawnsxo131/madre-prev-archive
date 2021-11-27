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

  protected calcMaxOfNumber(value: number) {
    const list = [
      0, 100, 1000, 5000, 10000, 50000, 100000, 500000, 1000000, 5000000,
      10000000, 50000000, 100000000, 1000000000, 5000000000, 10000000000,
      50000000000, 100000000000, 500000000000,
    ];
    for (let i = 0; i < list.length; i++) {
      if (value <= list[i]) {
        return list[i];
      }
    }
    return 0;
  }
}
