import type { ButtonHTMLAttributes } from 'react';

import { block } from './Button.css';

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  size: 'small' | 'medium' | 'large' | 'responsive';
};

export function Button({ children, ...props }: ButtonProps) {
  return (
    <button className={block} {...props}>
      {children}
    </button>
  );
}
