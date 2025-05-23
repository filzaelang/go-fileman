import { RiCloseLine } from "react-icons/ri";
import { useForm } from "@inertiajs/react";
import { useState, useEffect } from "react";

const ModalEditMenu = ({ setIsMEditOpen, onSubmit, id }) => {
  const [loading, setLoading] = useState(true);
  const { data, setData } = useForm({
    id: null,
    name: "",
    uri: "",
    parent_id: null,
  });

  useEffect(() => {
    fetch(`/api/menus/${id}`)
      .then((res) => res.json())
      .then((menu) => {
        setData({
          id: menu.id,
          name: menu.name,
          uri: menu.uri,
          parent_id: menu.parent_id,
        });
        setLoading(false);
      });
  }, [id, setData]);

  if (loading) return null;

  return (
    <>
      {/* Overlay */}
      <div
        className="fixed inset-0 bg-white/30 bg-opacity-20 z-40"
        onClick={(e) => {
          e.stopPropagation();
          setIsMEditOpen(false);
        }}
      />
      <div onClick={(e) => e.stopPropagation()}>
        {/* Modal Center */}
        <div className="fixed z-50 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[280px] bg-white rounded-xl shadow-lg">
          {/* Header */}
          <div className="flex justify-center items-center h-12 border-b border-gray-200 relative">
            <h5 className="text-gray-800 text-md font-semibold">Edit Menu</h5>
            <button
              className="absolute top-1 right-2 text-gray-500 hover:text-black"
              onClick={() => setIsMEditOpen(false)}
            >
              <RiCloseLine size={20} />
            </button>
          </div>

          {/* Content */}
          <div className="p-4">
            <form className="flex flex-col space-y-3">
              <label className="w-full text-gray-800 placeholder-gray-400">
                Nama Menu:
              </label>
              <input
                name="name"
                type="text"
                value={data.name}
                onChange={(e) => setData("name", e.target.value)}
                placeholder="Misal: Standard Procedure"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              <label className="w-full text-gray-800 placeholder-gray-400">
                Alamat Url:
              </label>
              <input
                type="text"
                value={data.uri}
                onChange={(e) => setData("uri", e.target.value)}
                placeholder="Misal: /standard-procedure"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
            </form>
          </div>

          {/* Actions */}
          <div className="flex justify-around pb-4">
            <button
              onClick={() => onSubmit(data)}
              className="bg-blue-500 text-white px-6 py-2 text-sm font-semibold rounded-lg hover:bg-blue-600 transition"
            >
              Submit
            </button>
            <button
              onClick={() => setIsMEditOpen(false)}
              className="bg-gray-100 text-gray-700 px-6 py-2 text-sm font-semibold rounded-lg hover:bg-gray-200 transition"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    </>
  );
};

export default ModalEditMenu;
