import React from "react"
import clsx from "clsx"

export interface IconProps extends React.SVGProps<SVGSVGElement>{}

const PencilIcon = React.forwardRef<SVGSVGElement, IconProps>(({ className, fill, stroke, ...props }, ref ) => (
  <svg ref={ref} {...props}
    className={clsx(className)}
    width="25"
    height="25"
    viewBox="0 0 512 512"
    fill={fill ?? "#FFFFFF"}
    stroke={stroke ?? "#FFFFFF"}
    xmlns="http://www.w3.org/2000/svg"
  >
    <g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <g> <g> <polygon points="24.735,388.224 0.308,512 123.538,487.026 "></polygon> </g> </g> <g> <g> <rect x="7.689" y="202.025" transform="matrix(0.7071 -0.7071 0.7071 0.7071 -101.9738 216.8106)" width="406.076" height="58.947"></rect> </g> </g> <g> <g> <path d="M500.185,67.086l-55.577-55.575c-15.349-15.347-40.228-15.347-55.578-0.001l-27.789,27.789l111.153,111.151l27.789-27.788 C515.529,107.314,515.529,82.43,500.185,67.086z"></path> </g> </g> <g> <g> <rect x="77.161" y="271.502" transform="matrix(0.7071 -0.7071 0.7071 0.7071 -130.7529 286.2839)" width="406.076" height="58.945"></rect> </g> </g> </g>
  </svg>
));

export default PencilIcon;
