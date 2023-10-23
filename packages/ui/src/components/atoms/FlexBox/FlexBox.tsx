import type { CSSProperties } from '@vanilla-extract/css';
import type { HTMLAttributes } from 'react';
import { Children } from 'react';

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
  gap?: {
    x?: number | string;
    y?: number | string;
    withFirstX?: boolean;
    withFirstY?: boolean;
  };
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
  gap,
  ...props
}: FlexBoxProps) {
  const containerStyle = Object.fromEntries(
    Object.entries({
      display: 'flex',
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
    }).filter(([_, val]) => !!val),
  );

  return (
    <div style={containerStyle} {...props}>
      {gap
        ? Children.map(children, (child, idx) => (
            <div
              style={{
                marginLeft: gap.withFirstX ? gap.x : idx !== 0 ? gap.x : 0,
                marginTop: gap.withFirstY ? gap.y : idx !== 0 ? gap.y : 0,
              }}
            >
              {child}
            </div>
          ))
        : children}
    </div>
  );
}
