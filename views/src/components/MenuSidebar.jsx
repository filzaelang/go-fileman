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

const staticMenus = [
  {
    headfolder: "Dashboard",
    uri: "/dashboard",
    icon: <FaHome className="icon-link" />,
  },
  {
    headfolder: "Profile",
    uri: "/profile",
    icon: <FaUser className="icon-link" />,
  },
  {
    headfolder: "Setup",
    icon: <FaGears className="icon-link" />,
    children: [
      {
        foldername: "Profile Authority",
        uri: "/setup/profile-authority",
        icon: <HiMiniUserGroup className="icon-link" />,
      },
      {
        foldername: "Folder Management",
        uri: "/setup/folder-management",
        icon: <FaFolderOpen className="icon-link" />,
      },
      {
        foldername: "Department Management",
        uri: "/setup/department-management",
        icon: <FaBuilding className="icon-link" />,
      },
      {
        foldername: "Check files",
        uri: "/setup/check-files",
        icon: <FaBriefcase className="icon-link" />,
      },
    ],
  },
  {
    headfolder: "Favorites",
    uri: "/favorites",
    icon: <FaStar className="icon-link" />,
  },
  {
    headfolder: "All Files",
    uri: "/allfiles",
    icon: <IoDocuments className="icon-link" />,
  },
];

function MenuSidebar({ menus = [], role = "super admin" }) {
  console.log("Ini menus", menus);
  return (
    <>
      {staticMenus.map((menu, index) => (
        <MenuItem key={index} menu={menu} isEditable={false} />
      ))}
      {menus !== null &&
        menus.map((menu, index) => (
          <MenuItem
            key={index}
            menu={menu}
            isEditable={role === "super admin"}
          />
        ))}
    </>
  );
}

export default MenuSidebar;
