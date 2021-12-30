import { format } from 'd3';

function formatNumberWithComma() {
  return format(',');
}

const D3FormatUtil = {
  formatNumberWithComma,
};

export default D3FormatUtil;
