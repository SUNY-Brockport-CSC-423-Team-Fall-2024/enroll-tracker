import React from "react";
import clsx from "clsx";

export interface IconProps extends React.SVGProps<SVGSVGElement> {}

const CapIcon = React.forwardRef<SVGSVGElement, IconProps>(
  ({ className, fill, stroke, ...props }, ref) => (
    <svg
      ref={ref}
      {...props}
      className={clsx(className)}
      width="25"
      height="25"
      viewBox="0 -4.83 31.876 31.876"
      fill={fill ?? "#FFFFFF"}
      xmlns="http://www.w3.org/2000/svg"
    >
      <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
      <g id="SVGRepo_tracerCarrier" strokeLinecap="round" strokeLinejoin="round"></g>
      <g id="SVGRepo_iconCarrier">
        {" "}
        <g transform="translate(-673.292 -327.728)">
          {" "}
          <path d="M689.741,329.778l9.241,2.448-9.242,2.7-9.82-2.7,9.821-2.449m.012-2.05a.478.478,0,0,0-.113.013L673.752,331.7a.465.465,0,0,0-.01.9l15.887,4.366a.467.467,0,0,0,.123.017.476.476,0,0,0,.13-.019l14.951-4.366a.465.465,0,0,0-.011-.9l-14.95-3.962a.479.479,0,0,0-.119-.015Z"></path>{" "}
          <path d="M696.013,349.95H682.63a3.932,3.932,0,0,1-4.124-3.7v-8.831a1,1,0,0,1,2,0v8.831a1.95,1.95,0,0,0,2.124,1.7h13.383a1.949,1.949,0,0,0,2.125-1.7v-8.831a1,1,0,0,1,2,0v8.831A3.932,3.932,0,0,1,696.013,349.95Z"></path>{" "}
          <path d="M674.292,341.16a1,1,0,0,1-1-1v-4.208a1,1,0,0,1,2,0v4.208A1,1,0,0,1,674.292,341.16Z"></path>{" "}
        </g>{" "}
      </g>
    </svg>
  ),
);

export default CapIcon;
