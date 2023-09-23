import { themePalette } from '../../styles';

interface CloseIconProps extends React.SVGProps<HTMLOrSVGElement> {}

function CloseIcon({ onClick, className }: CloseIconProps) {
  return (
    <svg
      stroke="currentColor"
      fill={themePalette.fill1}
      strokeWidth="0"
      viewBox="0 0 24 24"
      tabIndex={1}
      height="1em"
      width="1em"
      xmlns="http://www.w3.org/2000/svg"
      onClick={onClick}
      className={className}
    >
      <path
        d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"
        onClick={onClick}
      ></path>
    </svg>
  );
}

export default CloseIcon;
