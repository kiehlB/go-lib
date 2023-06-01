import clsx from "clsx";
import React from "react";
import { NavbarItem, NavbarItemProps } from "./index";

interface NavbarProps {
  primaryItems: NavbarItemProps[];
  secondaryItems?: NavbarItemProps[];
  className?: string;
  isDisabled?: boolean;
}

const Navbar = ({
  primaryItems,
  secondaryItems,
  className,
  isDisabled,
}: NavbarProps) => (
  <div className={clsx("w-[80%] mxl:w-[90%]", className)}>
    <ul>
      {primaryItems.map((itemProps) => (
        <li key={itemProps.text}>
          <NavbarItem {...itemProps} />
        </li>
      ))}
    </ul>
  </div>
);

export default Navbar;
