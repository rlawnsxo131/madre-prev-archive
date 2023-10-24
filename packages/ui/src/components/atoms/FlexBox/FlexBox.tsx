import { assignInlineVars } from '@vanilla-extract/dynamic';
import classNames from 'classnames';
import type { CSSProperties, HTMLAttributes } from 'react';

import { block, flexVar, sprinkles } from './FlexBox.css';

export type FlexBoxProps = HTMLAttributes<HTMLDivElement> & {
  flex?: CSSProperties['flex'];
  flexBasis?: CSSProperties['flexBasis'];
  flexDirection?: CSSProperties['flexDirection'];
  flexFlow?: CSSProperties['flexFlow'];
  flexGrow?: CSSProperties['flexGrow'];
  flexShrink?: CSSProperties['flexShrink'];
  justifyContent?: CSSProperties['justifyContent'];
  alignContent?: CSSProperties['alignContent'];
  alignItems?: CSSProperties['alignItems'];
  alignSelf?: CSSProperties['alignSelf'];
};

export function FlexBox({
  children,
  flex,
  flexBasis,
  flexDirection,
  flexFlow,
  flexGrow,
  flexShrink,
  justifyContent,
  alignContent,
  alignItems,
  alignSelf,
  className,
  ...props
}: FlexBoxProps) {
  const containerStyle = Object.fromEntries(
    Object.entries({
      display: 'flex',
      flexBasis,
      flexDirection,
      flexFlow,
      flexGrow,
      flexShrink,
      justifyContent,
      alignContent,
      alignItems,
      alignSelf,
    }).filter(([_, val]) => !!val),
  );

  return (
    <div
      className={classNames(block, sprinkles({ ...containerStyle }), className)}
      style={flex ? assignInlineVars({ [flexVar]: `${flex}` }) : undefined}
      {...props}
    >
      {children}
    </div>
  );
}
