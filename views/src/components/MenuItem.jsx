import { useState } from "react";
import { Link } from "@inertiajs/react";
import { FaChevronDown, FaFolder, FaChevronLeft } from "react-icons/fa";
import { FaPlus } from "react-icons/fa";
import { MdEdit } from "react-icons/md";
import { MdDelete } from "react-icons/md";
import { useForm } from "@inertiajs/react";

function MenuItem({ menu, level = 0, isEditable }) {
  const [open, setOpen] = useState(false);
  const hasChildren = menu.children && menu.children.length > 0;
  const { delete: destroy } = useForm();

  const remove = (id) => {
    if (confirm("Yakin mau hapus?")) {
      destroy(`/api/menus/${id}`, {
        preserveScroll: true,
      });
    }
  };

  return (
    <div className={`pl-${level * 4}`}>
      <div
        className="menu-link cursor-pointer flex flex-row items-center gap-2 justify-center"
        onClick={() => hasChildren && setOpen(!open)}
      >
        <div className="flex flex-row items-center justify-center sm:justify-between  sm:w-[100%] sm:me-3">
          <div className="">
            <div className="flex flex-row items-center gap-3">
              {menu.icon || <FaFolder className="icon-link" />}
              {menu.uri ? (
                <Link className="text-link" href={menu.uri}>
                  {menu.name}
                </Link>
              ) : (
                <span className="text-link">{menu.name}</span>
              )}
            </div>
          </div>
          <div className="flex flex-row gap-2">
            {hasChildren &&
              (open ? (
                <FaChevronDown size={12} className="text-link" />
              ) : (
                <FaChevronLeft size={12} className="text-link" />
              ))}
            {isEditable && (
              <MdDelete
                className="text-red-300 hidden sm:block"
                onClick={() => remove(menu.id)}
              />
            )}
            {isEditable && (
              <MdEdit
                className="text-gray-300 hidden sm:block"
                onClick={() => alert("Edit Menu")}
              />
            )}
            {isEditable && (
              <FaPlus
                className="text-gray-300 hidden sm:block"
                onClick={() => alert("Add Menu")}
              />
            )}
          </div>
        </div>
      </div>

      {hasChildren && open && (
        <div className="ml-4">
          {menu.children.map((child, idx) => (
            <MenuItem
              key={idx}
              menu={child}
              level={level + 1}
              isEditable={isEditable}
            />
          ))}
        </div>
      )}
    </div>
  );
}

export default MenuItem;
