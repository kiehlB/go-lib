"use client";

import clsx from "clsx";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React, { useRef, useEffect, useState } from "react";

export interface NavbarItemProps extends Pick<any, "to"> {
  text: string;
  to: string;
  sub?: string[];
}

export const NavbarItem = (props: NavbarItemProps) => {
  const pathname = usePathname();
  const isSelected =
    props.to === pathname ||
    pathname.startsWith(`${props.to}/`) ||
    (props.sub &&
      props.sub.some(
        (path) => pathname === path || pathname.startsWith(`${path}/`)
      ));

  const [iconAndTextWidth, setIconAndTextWidth] = useState<number>(0);
  const [textHeight, setTextHeight] = useState<number>(0);
  const iconAndTextRef = useRef<HTMLSpanElement>(null);
  const textRef = useRef<HTMLSpanElement>(null);

  useEffect(() => {
    if (iconAndTextRef.current) {
      const newWidth = iconAndTextRef.current.getBoundingClientRect().width;
      setIconAndTextWidth(newWidth);
    }
    if (textRef.current) {
      const newHeight = textRef.current.getBoundingClientRect().height;
      setTextHeight(newHeight);
    }
  }, [props.text]);

  return (
    <Link
      className={clsx(
        "underlined hover:text-team-current focus:text-team-current relative my-[0.5rem] flex items-center whitespace-nowrap py-[0.5rem] text-lg font-semibold transition-all focus:outline-none dark:text-[#e4e5e7]",
        {
          "active rounded text-[#212529] hover:pl-2 dark:bg-[#54565F33] dark:hover:bg-[#53525280]":
            isSelected,
          "text-[#495057] transition-all hover:pl-2": !isSelected,
        }
      )}
      href={props.to}
    >
      <span className={clsx("relative flex items-center")} ref={iconAndTextRef}>
        <span className='ml-2' ref={textRef}>
          {props.text}
        </span>
        {isSelected && (
          <span
            className='absolute'
            style={{
              backgroundColor: "#FFF23080",
              height: "12px",
              width: iconAndTextWidth,
              top: textHeight / 2 + 2,
              zIndex: -1,
            }}
          />
        )}
      </span>
    </Link>
  );
};
