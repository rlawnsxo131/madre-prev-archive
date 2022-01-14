import getRandomIntInclusive from './getRandomIntInclusive';

export default function getRandomColors(maxLength: number) {
  const result: string[] = [];
  const colors = [
    '#adb5bd',
    '#868e96',
    '#495057',
    '#343a40',
    '#212529',
    '#ff6b6b',
    '#fa5252',
    '#f03e3e',
    '#e03131',
    '#c92a2a',
    '#f06595',
    '#e64980',
    '#d6336c',
    '#c2255c',
    '#a61e4d',
    '#cc5de8',
    '#be4bdb',
    '#ae3ec9',
    '#9c36b5',
    '#862e9c',
    '#845ef7',
    '#7950f2',
    '#7048e8',
    '#6741d9',
    '#5f3dc4',
    '#5c7cfa',
    '#4c6ef5',
    '#4263eb',
    '#3b5bdb',
    '#364fc7',
    '#339af0',
    '#228be6',
    '#1c7ed6',
    '#1971c2',
    '#1864ab',
    '#22b8cf',
    '#15aabf',
    '#1098ad',
    '#0c8599',
    '#0b7285',
    '#20c997',
    '#12b886',
    '#0ca678',
    '#099268',
    '#087f5b',
    '#51cf66',
    '#40c057',
    '#37b24d',
    '#2f9e44',
    '#2b8a3e',
    '#94d82d',
    '#82c91e',
    '#74b816',
    '#66a80f',
    '#5c940d',
    '#fcc419',
    '#fab005',
    '#f59f00',
    '#f08c00',
    '#e67700',
    '#ff922b',
    '#fd7e14',
    '#f76707',
    '#e8590c',
    '#d9480f',
  ];
  const min = 0;
  const max = colors.length - 1;
  while (result.length < maxLength) {
    const randomNumber = getRandomIntInclusive(min, max);
    if (!result.includes(colors[randomNumber])) {
      result.push(colors[randomNumber]);
    }
  }
  return result;
}
