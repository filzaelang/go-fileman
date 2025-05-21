import React, { useState } from "react";
import { Link } from "@inertiajs/react";
import { FaChevronRight, FaChevronDown, FaFolder } from "react-icons/fa";

function MenuItem({ menu, level = 0 }) {
  const [open, setOpen] = useState(false);
  const hasChildren = menu.children && menu.children.length > 0;

  return (
    <div className={`pl-${level * 4}`}>
      <div
        className="menu-link cursor-pointer flex items-center gap-2"
        onClick={() => hasChildren && setOpen(!open)}
      >
        {menu.icon || <FaFolder className="icon-link" />}
        {menu.uri ? (
          <Link className="text-link" href={menu.uri}>
            {menu.name}
          </Link>
        ) : (
          <span className="text-link">{menu.name}</span>
        )}
        {hasChildren &&
          (open ? (
            <FaChevronDown size={12} className="text-link" />
          ) : (
            <FaChevronRight size={12} className="text-link" />
          ))}
      </div>

      {hasChildren && open && (
        <div className="ml-4">
          {menu.children.map((child, idx) => (
            <MenuItem key={idx} menu={child} level={level + 1} />
          ))}
        </div>
      )}
    </div>
  );
}

export default MenuItem;
