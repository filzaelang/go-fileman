import { FaHome } from "react-icons/fa";
import { FaUser } from "react-icons/fa6";
import { FaGears } from "react-icons/fa6";
import { FaStar } from "react-icons/fa6";
import { IoDocuments } from "react-icons/io5";
import MenuItem from "./MenuItem";
import { FaBuilding } from "react-icons/fa6";
import { FaBriefcase } from "react-icons/fa";
import { FaFolderOpen } from "react-icons/fa6";
import { HiMiniUserGroup } from "react-icons/hi2";

const menus = [
  {
    name: "Dashboard",
    uri: "/dashboard",
    icon: <FaHome className="icon-link" />,
  },
  {
    name: "Profile",
    uri: "/profile",
    icon: <FaUser className="icon-link" />,
  },
  {
    name: "Setup",
    icon: <FaGears className="icon-link" />,
    children: [
      {
        name: "Profile Authority",
        uri: "/setup/profile-authority",
        icon: <HiMiniUserGroup className="icon-link" />,
      },
      {
        name: "Folder Management",
        uri: "/setup/folder-management",
        icon: <FaFolderOpen className="icon-link" />,
      },
      {
        name: "Department Management",
        uri: "/setup/department-management",
        icon: <FaBuilding className="icon-link" />,
      },
      {
        name: "Check files",
        uri: "/setup/check-files",
        icon: <FaBriefcase className="icon-link" />,
      },
    ],
  },
  {
    name: "Favorites",
    uri: "/favorites",
    icon: <FaStar className="icon-link" />,
  },
  {
    name: "All Files",
    uri: "/allfiles",
    icon: <IoDocuments className="icon-link" />,
  },
  {
    name: "Dummy",
    children: [
      {
        name: "User Setup",
        uri: "/setup/user",
      },
      {
        name: "Roles",
        uri: "/setup/roles",
        children: [
          { name: "Admin", uri: "/setup/roles/admin" },
          { name: "User", uri: "/setup/roles/user" },
        ],
      },
    ],
  },
];

function MenuSidebar() {
  return (
    <>
      {menus.map((menu, index) => (
        <MenuItem key={index} menu={menu} />
      ))}
    </>
  );
}

export default MenuSidebar;
