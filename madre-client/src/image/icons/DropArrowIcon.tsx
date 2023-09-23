interface DropArrowIconProps extends React.SVGProps<HTMLOrSVGElement> {}

function DropArrowIcon(props: DropArrowIconProps) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
    >
      <path d="M12 21l-12-18h24z" />
    </svg>
  );
}

export default DropArrowIcon;
