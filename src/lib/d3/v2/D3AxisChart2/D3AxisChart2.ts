import D3Common2 from '../D3Common2';
import { D3SelectionSVG } from '../D3Common2/D3Common2Types';
import { D3AxisChartConstructorParams, D3Margin } from './D3AxisChart2Types';

export default class D3AxisChart2 extends D3Common2 {
  private svg: D3SelectionSVG;
  private margin: D3Margin;

  constructor({
    container,
    width,
    height,
    className = '',
    margin,
    data,
  }: D3AxisChartConstructorParams) {
    super();
    this.svg = this.appendSvg({
      container,
      width,
      height,
      className,
    });
    this.margin = margin;
  }
}
