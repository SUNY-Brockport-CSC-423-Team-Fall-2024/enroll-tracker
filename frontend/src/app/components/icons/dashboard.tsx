import React from "react";
import clsx from "clsx";

export interface IconProps extends React.SVGProps<SVGSVGElement> {}

const DashboardIcon = React.forwardRef<SVGSVGElement, IconProps>(
  ({ className, fill, stroke, ...props }, ref) => (
    <svg
      ref={ref}
      {...props}
      className={clsx(className)}
      width="25"
      height="25"
      viewBox="0 0 24 24"
      fill={fill ?? "none"}
      xmlns="http://www.w3.org/2000/svg"
    >
      <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
      <g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round"></g>
      <g id="SVGRepo_iconCarrier">
        {" "}
        <path
          d="M3 8.976C3 4.05476 4.05476 3 8.976 3H15.024C19.9452 3 21 4.05476 21 8.976V15.024C21 19.9452 19.9452 21 15.024 21H8.976C4.05476 21 3 19.9452 3 15.024V8.976Z"
          stroke={stroke}
          strokeWidth="2"
        ></path>{" "}
        <path
          d="M21 9L3 9"
          stroke={stroke}
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        ></path>{" "}
        <path
          d="M9 21L9 9"
          stroke={stroke}
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        ></path>{" "}
      </g>
    </svg>
  ),
);

export default DashboardIcon;
