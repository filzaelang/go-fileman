import React from "react";
import { RiCloseLine } from "react-icons/ri";
import { useForm } from "@inertiajs/react";

export default function ModalUpload({ setIsMUploadOpen, onSubmit, id }) {
  const { data, setData } = useForm({
    id: null,
    name: "",
    uri: "",
    parent_id: id,
  });

  return (
    <>
      {/* Overlay */}
      <div
        className="fixed inset-0 bg-white/30 bg-opacity-20 z-40"
        onClick={(e) => {
          e.stopPropagation();
          setIsMUploadOpen(false);
        }}
      />
      <div onClick={(e) => e.stopPropagation()}>
        {/* Modal Center */}
        <div className="fixed z-50 top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[280px] bg-white rounded-xl shadow-lg">
          {/* Header */}
          <div className="flex justify-center items-center h-12 border-b border-gray-200 relative">
            <h5 className="text-gray-800 text-md font-semibold">Upload File</h5>
            <button
              className="absolute top-1 right-2 text-gray-500 hover:text-black"
              onClick={() => setIsMUploadOpen(false)}
            >
              <RiCloseLine size={20} />
            </button>
          </div>

          {/* Content */}
          <div className="p-4">
            <form className="flex flex-col space-y-3">
              <label className="w-full text-gray-800 placeholder-gray-400">
                Document Number:
              </label>
              <input
                name="name"
                type="text"
                value={data.name}
                onChange={(e) => setData("name", e.target.value)}
                placeholder="No Dokumen"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              <label className="w-full text-gray-800 placeholder-gray-400">
                Document Number:
              </label>
              <input
                name="name"
                type="text"
                value={data.name}
                onChange={(e) => setData("name", e.target.value)}
                placeholder="No Dokumen"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              <label className="w-full text-gray-800 placeholder-gray-400">
                Document Name:
              </label>
              <input
                name="name"
                type="text"
                value={data.name}
                onChange={(e) => setData("name", e.target.value)}
                placeholder="Nama Dokumen"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              <label className="w-full text-gray-800 placeholder-gray-400">
                Revision Number:
              </label>
              <input
                name="name"
                type="text"
                value={data.name}
                onChange={(e) => setData("name", e.target.value)}
                placeholder="No Revisi"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              <label className="w-full text-gray-800 placeholder-gray-400">
                Revision Date:
              </label>
              <input type="date" id="birthday" name="birthday"></input>
              <input type="file" id="myfile" name="myfile"></input>
              <button className="bg-blue-400" onClick={() => onSubmit(data)}>
                Upload
              </button>
              <button
                className="bg-gray-400"
                onClick={() => setIsMUploadOpen(false)}
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
