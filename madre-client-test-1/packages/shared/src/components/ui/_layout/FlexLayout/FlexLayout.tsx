import classNames from 'classnames';
import type {
  ComponentPropsWithoutRef,
  ComponentPropsWithRef,
  CSSProperties,
  ElementType,
  ReactNode,
} from 'react';
import { forwardRef } from 'react';

import { sprinkles } from './FlexLayout.css';

export type FlexLayoutProps<E extends ElementType> =
  ComponentPropsWithoutRef<E> & {
    as?: E;
    children?: ReactNode;
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

export type FlexLayoutComponent = <E extends ElementType = 'div'>(
  Props: FlexLayoutProps<E> & { ref?: ComponentPropsWithRef<E>['ref'] },
) => ReactNode;

export const FlexLayout: FlexLayoutComponent = forwardRef(function <
  E extends ElementType,
>(
  {
    as,
    children,
    className,
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
  }: FlexLayoutProps<E>,
  ref?: ComponentPropsWithRef<E>['ref'],
) {
  const Element = as || 'div';
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
    <Element
      ref={ref}
      className={classNames(sprinkles({ ...flexStyles }), className)}
      {...props}
    >
      {children}
    </Element>
  );
});
