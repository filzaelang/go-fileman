import Layout from "../../layouts/Layout";
import { usePage } from "@inertiajs/react";
import { router } from "@inertiajs/react";
import ModalUpload from "../../components/Modal/ModalUpload";
import { useState } from "react";

function Index({ phrase, items }) {
  const { url } = usePage();
  const [isMUploadOpen, setIsMUploadOpen] = useState(false);

  const upload = (data) => {
    setIsMUploadOpen(false);
    router.post("/api/files", data, {
      onSuccess: () => router.visit("/"),
    });
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
              onClick={(e) => {
                e.stopPropagation();
                setIsMUploadOpen(true);
              }}
            >
              Upload File
            </button>
            <button className="bg-gray-100 p-2 rounded-md hover:bg-gray-200">
              Move File
            </button>
          </div>
        </div>
        {isMUploadOpen && (
          <ModalUpload setIsMUploadOpen={setIsMUploadOpen} onSubmit={upload} />
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
          <div class="relative flex flex-col w-full h-full overflow-scroll text-gray-700 bg-white shadow-md rounded-xl bg-clip-border">
            <table class="w-full text-left table-auto min-w-max">
              <thead>
                <tr>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Checklist
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      No.
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Document Number
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Old Document Number
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Document Name
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Revision Number
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Revision Date
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Fav
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Visible
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Download
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Detail
                    </p>
                  </th>
                  <th class="p-4 border-b border-blue-gray-100 bg-blue-gray-50">
                    <p class="block font-sans text-sm antialiased font-normal leading-none text-blue-gray-900 opacity-70">
                      Delete
                    </p>
                  </th>
                </tr>
              </thead>
              <tbody>
                {items.map((item, index) => (
                  <tr>
                    <td class="p-4 border-b border-blue-gray-50">
                      <p class="block font-sans text-sm antialiased font-normal leading-normal text-blue-gray-900">
                        V
                      </p>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <p class="block font-sans text-sm antialiased font-normal leading-normal text-blue-gray-900">
                        {index + 1}
                      </p>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <p class="block font-sans text-sm antialiased font-normal leading-normal text-blue-gray-900">
                        {item.filenumber}
                      </p>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <a
                        href="#"
                        class="block font-sans text-sm antialiased font-medium leading-normal text-blue-gray-900"
                      >
                        {item.fileoldnumber}
                      </a>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <p>{item.filename}</p>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <p>{item.filerevnumber}</p>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <p>{item.filerevdate}</p>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <p>favorites</p>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <p>{item.filevisible}</p>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <a
                        href={`/api/files/${item.fileoid}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="bg-green-500 py-1 px-2 text-sm rounded-md text-white hover:bg-green-700 focus:ring-gray-700"
                      >
                        Download PDF
                      </a>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      <button className="bg-blue-500 py-1 px-2 text-sm rounded-md text-white hover:bg-blue-700">
                        Edit
                      </button>
                    </td>
                    <td class="p-4 border-b border-blue-gray-50">
                      {/* <button
                        className="bg-red-500 py-1 px-2 text-sm rounded-md text-white hover:bg-red-700"
                        onClick={deleteFile(item.fileoid)}
                      >
                        Delete
                      </button> */}
                    </td>
                  </tr>
                ))}
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
