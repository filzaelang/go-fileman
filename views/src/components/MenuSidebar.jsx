import { Link } from "@inertiajs/react";
import { FaHome } from "react-icons/fa";
import { FaUser } from "react-icons/fa6";
import { FaGears } from "react-icons/fa6";
import { FaStar } from "react-icons/fa6";
import { IoDocuments } from "react-icons/io5";
import MenuItem from "./MenuItem";

function MenuSidebar() {
  return (
    <>
      <div className="flex flex-col">
        <div className="menu-link">
          <FaHome className="icon-link" />
          <Link className="text-link" href="/dashboard">
            Dashboard
          </Link>
        </div>
        <div className="menu-link">
          <FaUser className="icon-link" />
          <Link className="text-link" href="/profile">
            Profile
          </Link>
          <span className="tooltip">Profile</span>
        </div>
        <div className="menu-link">
          <FaGears className="icon-link" />
          <Link className="text-link" href="/setup">
            Setup
          </Link>
        </div>
        <div className="menu-link">
          <FaStar className="icon-link" />
          <Link className="text-link" href="/favorites">
            Favorites
          </Link>
        </div>
        <div className="menu-link">
          <IoDocuments className="icon-link" />
          <Link className="text-link" href="/allfiles">
            All Files
          </Link>
        </div>
        <MenuItem
          name={"Dummy Quality Control"}
          uri={"/dummy-quality-control"}
        />
        <MenuItem
          name={"Dummy Standard Procedure"}
          uri={"/dummy-standard-procedure"}
        />
        <MenuItem
          name={"Dummy Standard Instruction"}
          uri={"/dummy-standard-instruction"}
        />
      </div>
    </>
  );
}

export default MenuSidebar;
