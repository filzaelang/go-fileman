import React from "react";
import "./ModalDeleteMenu.css";
import { RiCloseLine } from "react-icons/ri";

const ModalDeleteMenu = ({ setIsMDelOpen, onDelete, id }) => {
  return (
    <>
      <div className={"darkBG"} onClick={() => setIsMDelOpen(false)} />
      <div className={"centered"}>
        <div className={"modal"}>
          <div className={"modalHeader"}>
            <h5 className={"heading"}>PERINGATAN</h5>
          </div>
          <button className={"closeBtn"} onClick={() => setIsMDelOpen(false)}>
            <RiCloseLine style={{ marginBottom: "-3px" }} />
          </button>
          <div className={"modalContent"}>
            Apakah anda yakin akan menghapus menu ini beserta seluruh sub menu
            nya ?
          </div>
          <div className={"modalActions mt-2"}>
            <div className={"actionsContainer"}>
              <button className={"deleteBtn"} onClick={() => onDelete(id)}>
                Ya
              </button>
              <button
                className={"cancelBtn"}
                onClick={() => setIsMDelOpen(false)}
              >
                Tidak
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default ModalDeleteMenu;
