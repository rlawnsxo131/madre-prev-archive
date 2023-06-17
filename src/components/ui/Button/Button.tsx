import type { ButtonHTMLAttributes, PropsWithChildren } from 'react';

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & PropsWithChildren;

export function Button({ children, ...props }: ButtonProps) {
  return <button {...props}>{children}</button>;
}
