import { assignInlineVars } from '@vanilla-extract/dynamic';
import classNames from 'classnames';
import type { CSSProperties, HTMLAttributes, PropsWithoutRef } from 'react';
import { forwardRef } from 'react';

import { block, flexVar, sprinkles } from './FlexLayout.css';

export type FlexLayoutProps = PropsWithoutRef<
  HTMLAttributes<HTMLDivElement> & {
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
  }
>;

export const FlexLayout = forwardRef<HTMLDivElement, FlexLayoutProps>(function (
  {
    children,
    className,
    flex = '0 1 auto',
    flexBasis,
    flexDirection,
    flexFlow,
    flexGrow,
    flexShrink,
    justifyContent,
    alignContent,
    alignItems,
    alignSelf,
    ...props
  },
  ref,
) {
  const flexStyles = Object.fromEntries(
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
      ref={ref}
      className={classNames(block, sprinkles({ ...flexStyles }), className)}
      style={assignInlineVars({ [flexVar]: `${flex}` })}
      {...props}
    >
      {children}
    </div>
  );
});
