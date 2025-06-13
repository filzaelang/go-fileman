import { RiCloseLine } from "react-icons/ri";
import { useForm } from "@inertiajs/react";

export default function ModalUpload({
  setIsMUploadOpen,
  onSubmit,
  folderoid,
  divoid,
  deptoid,
}) {
  const { data, setData } = useForm({
    document_number: "",
    document_name: "",
    revision_number: "",
    revision_date: "",
    folderoid: folderoid,
    divoid: divoid,
    deptoid: deptoid,
    file: null,
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
            <form
              className="flex flex-col space-y-3"
              onSubmit={(e) => {
                e.preventDefault();
                onSubmit(data);
              }}
            >
              <label className="w-full text-gray-800 placeholder-gray-400">
                Document Number:
              </label>
              <input
                name="document_number"
                type="text"
                value={data.document_number}
                onChange={(e) => setData("document_number", e.target.value)}
                placeholder="No Dokumen"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              <label className="w-full text-gray-800 placeholder-gray-400">
                Document Name:
              </label>
              <input
                name="document_name"
                type="text"
                value={data.document_name}
                onChange={(e) => setData("document_name", e.target.value)}
                placeholder="Nama Dokumen"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              <label className="w-full text-gray-800 placeholder-gray-400">
                Revision Number:
              </label>
              <input
                name="revision_number"
                type="text"
                value={data.revision_number}
                onChange={(e) => setData("revision_number", e.target.value)}
                placeholder="No Revisi"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              <label className="w-full text-gray-800 placeholder-gray-400">
                Revision Date:
              </label>
              <input
                type="date"
                id="revision_date"
                name="revision_date"
                onChange={(e) => setData("revision_date", e.target.value)}
              ></input>
              <input
                type="file"
                id="file"
                name="file"
                onChange={(e) => setData("file", e.target.files[0])}
              ></input>
              <button className="bg-blue-400" type="submit">
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
