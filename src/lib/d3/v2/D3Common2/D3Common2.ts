import { select } from 'd3';
import { AppendSVGParams } from './D3Common2Types';

export default class D3Common {
  private readonly displayMaxNumberArray = [
    0, 100, 1_000, 2_000, 3_000, 5_000, 8_000, 10_000, 30_000, 50_000, 100_000,
    500_000, 1_000_000, 5_000_000, 10_000_000, 50_000_000, 100_000_000,
    1_000_000_000, 5_000_000_000, 10_000_000_000, 50_000_000_000,
    100_000_000_000, 500_000_000_000,
  ];

  protected appendSVG({
    container,
    width,
    height,
    className = '',
  }: AppendSVGParams) {
    return (
      select(container)
        .append('svg')
        // .attr('width', '100%')
        // .attr('height', '100%')
        .attr('width', width)
        .attr('height', height)
        .attr('viewBox', `0, 0, ${width}, ${height}`)
        .attr('class', className)
    );
  }

  protected generateClassList(count: number) {}

  protected calcMaxOfNumber(value: number) {
    for (let i = 0; i < this.displayMaxNumberArray.length; i++) {
      if (value <= this.displayMaxNumberArray[i]) {
        return this.displayMaxNumberArray[i];
      }
    }
    return 900_719_925_474;
  }
}
