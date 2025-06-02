import { RiCloseLine } from "react-icons/ri";
import { useForm } from "@inertiajs/react";

export default function ModalMoveFile({ setIsMMoveFileOpen, onSubmit, menus }) {
  const { data, setData } = useForm({
    id: null,
    document_number: "",
    document_name: "",
    revision_number: "",
    revision_date: "",
    file: null,
  });

  return (
    <>
      {/* Overlay */}
      <div
        className="fixed inset-0 bg-white/30 bg-opacity-20 z-40"
        onClick={(e) => {
          e.stopPropagation();
          setIsMMoveFileOpen(false);
        }}
      />
      <div onClick={(e) => e.stopPropagation()}>
        {/* Modal Center */}
        <div className="fixed z-50 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[280px] bg-white rounded-xl shadow-lg">
          {/* Header */}
          <div className="flex justify-center items-center h-12 border-b border-gray-200 relative">
            <h5 className="text-gray-800 text-md font-semibold">Move Files</h5>
            <button
              className="absolute top-1 right-2 text-gray-500 hover:text-black"
              onClick={() => setIsMMoveFileOpen(false)}
            >
              <RiCloseLine size={20} />
            </button>
          </div>

          {/* Content */}
          <div className="p-4">
            <form
              className="flex flex-col space-y-3"
              onSubmit={(e) => {
                e.preventDefault();
                onSubmit(data);
              }}
            >
              <label className="w-full text-gray-800 placeholder-gray-400">
                Head folder
              </label>
              <select name="cars" id="cars">
                {menus.map((menu) => (
                  <option value={menu.headfolder}>{menu.headfolder}</option>
                ))}
              </select>
              <button className="bg-blue-400" type="submit">
                Move
              </button>
              <button
                className="bg-gray-400"
                onClick={() => setIsMMoveFileOpen(false)}
              >
                Close
              </button>
            </form>
          </div>
        </div>
      </div>
    </>
  );
}
