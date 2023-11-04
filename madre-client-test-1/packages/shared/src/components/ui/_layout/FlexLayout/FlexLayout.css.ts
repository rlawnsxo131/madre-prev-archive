import { createSprinkles, defineProperties } from '@vanilla-extract/sprinkles';

const properties = defineProperties({
  properties: {
    display: ['flex'],
    flexDirection: [
      'inherit',
      'row',
      '-moz-initial',
      'initial',
      'revert',
      'revert-layer',
      'unset',
      'column',
      'column-reverse',
      'row-reverse',
    ],
    flexFlow: [
      'row nowrap',
      'row wrap',
      'row wrap-reverse',
      'row-reverse nowrap',
      'row-reverse wrap',
      'row-reverse wrap-reverse',
      'column nowrap',
      'column wrap',
      'column wrap-reverse',
      'column-reverse nowrap',
      'column-reverse wrap',
      'column-reverse wrap-reverse',
    ],
    justifyContent: [
      'start',
      'center',
      'flex-start',
      'flex-end',
      'left',
      'right',
      'stretch',
      'space-between',
      'space-around',
      'space-evenly',
    ],
    alignContent: [
      'stretch',
      'center',
      'flex-start',
      'flex-end',
      'space-between',
      'space-around',
    ],
    alignItems: ['stretch', 'center', 'flex-start', 'flex-end'],
    alignSelf: ['auto', 'stretch', 'center', 'flex-start', 'flex-end'],
  },
});

export const sprinkles = createSprinkles(properties);
