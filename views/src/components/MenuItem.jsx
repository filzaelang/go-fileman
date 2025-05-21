import React from "react";
import { FaFolder } from "react-icons/fa";
import { Link } from "@inertiajs/react";

function MenuItem({ name, uri }) {
  return (
    <>
      <div className="menu-link">
        <FaFolder className="icon-link" />
        <Link className="text-link" href={uri}>
          {name}
        </Link>
      </div>
    </>
  );
}

export default MenuItem;
