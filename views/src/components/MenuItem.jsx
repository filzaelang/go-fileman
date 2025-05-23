import { useState } from "react";
import { Link } from "@inertiajs/react";
import { FaChevronDown, FaFolder, FaChevronLeft } from "react-icons/fa";
import { FaPlus } from "react-icons/fa";
import { MdEdit } from "react-icons/md";
import { MdDelete } from "react-icons/md";
import { router } from "@inertiajs/react";
import ModalDeleteMenu from "./Modal/ModalDeleteMenu";
import ModalAddMenu from "./Modal/ModalAddMenu";
import ModalEditMenu from "./Modal/ModalEditMenu";

function MenuItem({ menu, level = 0, isEditable }) {
  const [open, setOpen] = useState(false);
  const hasChildren = menu.children && menu.children.length > 0;
  const [isMAddOpen, setIsMAddOpen] = useState(false);
  const [isMDelOpen, setIsMDelOpen] = useState(false);
  const [isMEditOpen, setIsMEditOpen] = useState(false);

  const add = (data) => {
    setIsMAddOpen(false);
    router.post("/api/menus", data, {
      onSuccess: () => router.visit("/"),
    });
  };

  const edit = (data) => {
    setIsMEditOpen(false);
    router.put(`api/menus/${data.id}`, data, {
      onSuccess: () => router.visit("/"),
    });
  };

  const remove = (id) => {
    setIsMDelOpen(false);
    router.delete(`/api/menus/${id}`, {
      onSuccess: () => router.visit("/"),
    });
  };

  return (
    <div className={`pl-${level * 4}`}>
      <div
        className="menu-link cursor-pointer flex flex-row items-center gap-2 justify-center"
        onClick={(e) => {
          e.stopPropagation();
          hasChildren && setOpen(!open);
        }}
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
                onClick={(e) => {
                  e.stopPropagation();
                  setIsMDelOpen(true);
                }}
              />
            )}
            {isMDelOpen && (
              <ModalDeleteMenu
                setIsMDelOpen={setIsMDelOpen}
                onDelete={remove}
                id={menu.id}
              />
            )}
            {isEditable && (
              <MdEdit
                className="text-gray-300 hidden sm:block"
                onClick={(e) => {
                  e.stopPropagation();
                  setIsMEditOpen(true);
                }}
              />
            )}
            {isMEditOpen && (
              <ModalEditMenu
                setIsMEditOpen={setIsMEditOpen}
                onSubmit={edit}
                id={menu.id}
              />
            )}
            {isEditable && (
              <FaPlus
                className="text-gray-300 hidden sm:block"
                onClick={(e) => {
                  e.stopPropagation();
                  setIsMAddOpen(true);
                }}
              />
            )}
            {isMAddOpen && (
              <ModalAddMenu
                setIsMAddOpen={setIsMAddOpen}
                onSubmit={add}
                id={menu.id}
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
