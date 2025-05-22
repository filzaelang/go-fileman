import { Link } from "@inertiajs/react";
import { GiHamburgerMenu } from "react-icons/gi";
import { GoDotFill } from "react-icons/go";
import { FaUserCircle } from "react-icons/fa";
import { MdOutlineLogout } from "react-icons/md";
import MenuSidebar from "../components/MenuSidebar";
import { FaPlus } from "react-icons/fa";

export default function Layout({ children, menus }) {
  return (
    <div className="flex flex-row min-h-screen">
      {/* Sidebar */}
      <div
        className={`sidebar transition-all duration-300 ease-in-out bg-gray-700 min-h-screen w-18 sm:w-64`}
      >
        <div className="h-10 bg-cyan-600 flex justify-center items-center">
          <p className="text-white font-bold hidden sm:block">FILE MANAGER</p>
        </div>
        <div className="flex flex-row py-3 justify-center items-center">
          <div className="sm:ms-3 w-1/4 flex flex-row">
            <p className="text-xl justify-center  sm:text-5xl text-white">
              <FaUserCircle />
            </p>
          </div>
          <div className="flex-col w-3/4 justify-start hidden sm:block">
            <p className="text-white font-bold">Administrator</p>
            <div className="flex flex-row items-center text-white text-sm">
              <GoDotFill /> Online
            </div>
          </div>
        </div>
        <div className="hidden bg-gray-900 sm:flex sm:flex-row sm:p-3 sm:justify-between">
          <div className="">
            <p className="text-sm text-gray-600">MENU</p>
          </div>
          <div className="">
            <FaPlus className="text-gray-300" />
          </div>
        </div>
        <MenuSidebar menus={menus} />
        <div className="bg-gray-900 p-3 hidden sm:block">
          <p className="text-sm text-gray-600">LOG OUT</p>
        </div>
        <div className="menu-link">
          <MdOutlineLogout className="text-blue-500" />
          <Link className="text-link">Logout</Link>
        </div>
      </div>
      <div className="flex-grow">
        {/* Main Content */}
        <div className="h-10 bg-cyan-500 flex items-center">
          <button
            className="p-3 text-white"
            // onClick={() => setIsSidebarOpen(!isSidebarOpen)}
            onClick={() => alert("Hello")}
          >
            <GiHamburgerMenu />
          </button>
        </div>
        <div className="ms-3 mt-7">
          <main>{children}</main>
        </div>
      </div>
    </div>
  );
}
