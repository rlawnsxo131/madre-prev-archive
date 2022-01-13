import { v4 } from 'uuid';
import { getRandomColors } from '../../utils';

export function generateUniqIdentifierValueAndColorArray(length: number) {
  const colors = getRandomColors(length);
  return Array.from({ length }).map((v, i) => ({
    key: v4(),
    color: colors[i],
  }));
}
