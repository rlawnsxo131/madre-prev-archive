import classNames from 'classnames';
import { type ButtonHTMLAttributes } from 'react';

import {
  block,
  radius as _radius,
  size as _size,
  theme as _theme,
} from './Button.css';

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  theme?: 'primary' | 'secondary' | 'warn';
  size?: 'small' | 'medium' | 'large' | 'responsive';
  radius?: 'none' | 'medium' | 'full';
  variant?: 'solid' | 'outline';
};

export function Button({
  children,
  theme = 'primary',
  size = 'medium',
  radius = 'medium',
  variant = 'solid',
  ...props
}: ButtonProps) {
  return (
    <button
      className={classNames(block, _theme[theme], _size[size], _radius[radius])}
      {...props}
    >
      {children}
    </button>
  );
}
