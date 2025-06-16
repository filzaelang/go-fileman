import { RiCloseLine } from "react-icons/ri";
import { useState, useEffect } from "react";

export default function ModalUpload({
  setIsMUploadOpen,
  onSubmit,
  folderoid,
  divoid,
  deptoid,
}) {
  const [formData, setFormData] = useState({
    document_number: "",
    document_name: "",
    revision_number: "",
    revision_date: "",
    folderoid: folderoid,
    divoid: divoid,
    deptoid: deptoid,
    file: null,
  });

  const [touched, setTouched] = useState({
    document_number: false,
    document_name: false,
    revision_number: false,
    revision_date: false,
  });

  const [formValid, setFormValid] = useState(false);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleFileChange = (e) => {
    if (e.target.files && e.target.files.length > 0) {
      setFormData((prev) => ({ ...prev, file: e.target.files[0] }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = new FormData();
    data.append("document_number", formData.document_number);
    data.append("document_name", formData.document_name);
    data.append("revision_number", formData.revision_number);
    data.append("revision_date", formData.revision_date);
    data.append("folderoid", formData.folderoid);
    data.append("divoid", formData.divoid);
    data.append("deptoid", formData.deptoid);
    data.append("file", formData.file);

    onSubmit(data);
  };

  useEffect(() => {
    const isValid =
      formData.document_number.trim() !== "" &&
      formData.document_name.trim().length >= 3 &&
      formData.revision_date.trim() !== "" &&
      formData.file !== null;

    setFormValid(isValid);
  }, [formData]);

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
            <form className="flex flex-col space-y-3" onSubmit={handleSubmit}>
              <label className="w-full text-gray-800 placeholder-gray-400">
                Document Number<span className="text-red-500 text-xs">*</span>
              </label>
              <input
                name="document_number"
                type="text"
                value={formData.document_number}
                onBlur={() =>
                  setTouched((prev) => ({ ...prev, document_number: true }))
                }
                onChange={handleChange}
                placeholder="No Dokumen"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              {touched.document_number &&
                formData.document_number.trim() === "" && (
                  <p className="text-red-500 text-xs">
                    Nomor dokumen tidak boleh kosong
                  </p>
                )}
              <label className="w-full text-gray-800 placeholder-gray-400">
                Document Name<span className="text-red-500 text-xs">*</span>
              </label>
              <input
                name="document_name"
                type="text"
                value={formData.document_name}
                onBlur={() =>
                  setTouched((prev) => ({ ...prev, document_name: true }))
                }
                onChange={handleChange}
                placeholder="Nama Dokumen"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              {touched.document_name && (
                <>
                  {formData.document_name.trim() === "" ? (
                    <p className="text-red-500 text-xs">
                      Nama dokumen tidak boleh kosong
                    </p>
                  ) : formData.document_name.trim().length < 3 ? (
                    <p className="text-red-500 text-xs">
                      Nama dokumen minimal 3 karakter
                    </p>
                  ) : null}
                </>
              )}
              <label className="w-full text-gray-800 placeholder-gray-400">
                Revision Number<span className="text-red-500 text-xs">*</span>
              </label>
              <input
                name="revision_number"
                type="number"
                value={formData.revision_number}
                onBlur={() =>
                  setTouched((prev) => ({ ...prev, revision_number: true }))
                }
                onChange={handleChange}
                placeholder="Isi (00) jika belum ada revisi"
                className="w-full px-3 py-2 text-gray-800 placeholder-gray-400 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
              />
              {touched.revision_number &&
                formData.revision_number.trim() === "" && (
                  <p className="text-red-500 text-xs">
                    Nomor revisi tidak boleh kosong
                  </p>
                )}
              <p className="text-red-500 text-xs">
                Format pengisian : 00, 01, 02, 10, 11, dst
              </p>
              <label className="w-full text-gray-800 placeholder-gray-400">
                Revision Date<span className="text-red-500 text-xs">*</span>
              </label>
              <input
                type="date"
                id="revision_date"
                name="revision_date"
                onBlur={() =>
                  setTouched((prev) => ({ ...prev, revision_date: true }))
                }
                onChange={handleChange}
              ></input>
              {touched.revision_date &&
                formData.revision_date.trim() === "" && (
                  <p className="text-red-500 text-xs">
                    Tanggal revisi tidak boleh kosong
                  </p>
                )}
              <input
                type="file"
                id="file"
                name="file"
                onChange={handleFileChange}
              ></input>
              {formData.file === null && (
                <p className="text-red-500 text-xs">
                  File yang diupload tidak boleh kosong
                </p>
              )}
              <button
                type="submit"
                disabled={!formValid}
                className={`px-4 py-2 text-white rounded-md transition 
                  ${
                    formValid
                      ? "bg-blue-500 hover:bg-blue-600"
                      : "bg-gray-400 cursor-not-allowed"
                  }`}
              >
                Upload
              </button>
              <button
                className="bg-gray-400 hover:bg-gray-500"
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
