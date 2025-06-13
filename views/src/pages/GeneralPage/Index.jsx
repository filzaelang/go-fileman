import Layout from "../../layouts/Layout";
import { usePage } from "@inertiajs/react";
import { router } from "@inertiajs/react";
import { useState } from "react";
import ModalUpload from "../../components/Modal/ModalUpload";
import ModalMoveFile from "../../components/Modal/ModalMoveFile";
import axios from "axios";

function Index({ phrase, items, menus }) {
  const { url } = usePage();
  const parts = url.split("/");
  const folderoid = parts[2];
  const divoid = parts[3];
  const deptoid = parts[4];
  const [isMUploadOpen, setIsMUploadOpen] = useState(false);
  const [isMMoveFileOpen, setIsMMoveFileOpen] = useState(false);

  const upload = async (data) => {
    setIsMUploadOpen(false);

    const formData = new FormData();
    for (const key in data) {
      formData.append(key, data[key]);
    }

    try {
      const response = await axios.post("/api/files", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });

      const res = response.data;

      alert(res.message);
      router.visit(res.redirect);
    } catch (error) {
      alert(
        "Upload gagal: " + (error.response?.data?.message || error.message)
      );
    }
  };

  const moveFile = () => {
    setIsMMoveFileOpen(false);
    alert("Oh begitu saja");
  };

  const deleteFile = (id) => {
    router.delete(`/api/files/${id}`, {
      onSuccess: () => router.visit("/"),
    });
  };

  return (
    <>
      <div>
        <h1 className="text-black text-2xl">{url}</h1>
      </div>
      <div className="mt-5">
        <p className="text-2xl">{phrase}</p>
        <div className="mt-2">
          <div className="flex flex-row gap-2">
            <button
              className="bg-blue-400 p-2 rounded-md text-white hover:bg-blue-500"
              onClick={() => setIsMUploadOpen(true)}
            >
              Upload File
            </button>
            <button
              className="bg-gray-100 p-2 rounded-md hover:bg-gray-200"
              onClick={() => setIsMMoveFileOpen(true)}
            >
              Move File
            </button>
          </div>
        </div>
        {isMUploadOpen && (
          <ModalUpload
            setIsMUploadOpen={setIsMUploadOpen}
            onSubmit={upload}
            folderoid={folderoid}
            divoid={divoid}
            deptoid={deptoid}
          />
        )}
        {isMMoveFileOpen && (
          <ModalMoveFile
            menus={menus}
            setIsMMoveFileOpen={setIsMMoveFileOpen}
            onSubmit={moveFile}
          />
        )}
        <div>
          <div className="mt-2">
            <span>Show</span>
            <select
              name="pagination"
              id="volume"
              className="bg-gray-100 p-1 mx-2 rounded-md border-1"
            >
              <option>10</option>
              <option>25</option>
              <option>50</option>
              <option>100</option>
            </select>
            <span>entries</span>
          </div>
        </div>
        <div>
          <div className="relative flex flex-col w-full h-full overflow-scroll text-gray-700 bg-white shadow-md rounded-xl bg-clip-border">
            <table className="w-full text-left table-auto min-w-max">
              <thead>
                <tr>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Checklist
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      No.
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Document Number
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Old Document Number
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Document Name
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Revision Number
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Revision Date
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Fav
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Visible
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Download
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Detail
                    </p>
                  </th>
                  <th className="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p className="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Delete
                    </p>
                  </th>
                </tr>
              </thead>
              <tbody>
                {items !== null ? (
                  items.map((item, index) => (
                    <tr key={index}>
                      <td className="p-4 border-b border-blue-gray-50">
                        <p className="block font-sans text-sm antialiased font-normal leading-normal text-blue-gray-900">
                          V
                        </p>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <p className="block font-sans text-sm antialiased font-normal leading-normal text-blue-gray-900">
                          {index + 1}
                        </p>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <p className="block font-sans text-sm antialiased font-normal leading-normal text-blue-gray-900">
                          {item.filenumber}
                        </p>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <a
                          href="#"
                          className="block font-sans text-sm antialiased font-medium leading-normal text-blue-gray-900"
                        >
                          {item.fileoldnumber}
                        </a>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <p>{item.filename}</p>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <p>{item.filerevnumber}</p>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <p>{item.filerevdate}</p>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <p>favorites</p>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <p>{item.filevisible}</p>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <a
                          href={`/api/files/${item.fileoid}`}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="bg-green-500 py-1 px-2 text-sm rounded-md text-white hover:bg-green-700 focus:ring-gray-700"
                        >
                          Download
                        </a>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <button className="bg-blue-500 py-1 px-2 text-sm rounded-md text-white hover:bg-blue-700">
                          Edit
                        </button>
                      </td>
                      <td className="p-4 border-b border-blue-gray-50">
                        <button
                          className="bg-red-500 py-1 px-2 text-sm rounded-md text-white hover:bg-red-700"
                          onClick={(e) => {
                            e.stopPropagation();
                            deleteFile(item.fileoid);
                          }}
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  ))
                ) : (
                  <tr>
                    <td colSpan="13" className="text-center py-4 text-gray-500">
                      No data available
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </>
  );
}

Index.layout = (page) => {
  const props = page.props;
  return (
    <Layout menus={props.menus} role={props.role}>
      {page}
    </Layout>
  );
};

export default Index;
