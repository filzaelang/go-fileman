import { useForm, usePage, router } from "@inertiajs/react";
import MenuItem from "../../components/MenuItem";

export default function SetupMenu({ menus }) {
  const {
    data,
    setData,
    post,
    put,
    delete: destroy,
    reset,
  } = useForm({
    id: null,
    name: "",
    uri: "",
    parent_id: null,
  });

  const submit = () => {
    // alert(`${data.name} - ${data.uri} - ${data.parent_id}`);
    // post("/api/menus", {
    //   onSuccess: reset,
    // });
    if (data.id) {
      put(`/api/menus/${data.id}`, {
        onSuccess: reset,
      });
    } else {
      post("/api/menus", {
        onSuccess: reset,
      });
    }
  };

  const remove = (id) => {
    if (confirm("Yakin mau hapus?")) {
      destroy(`/api/menus/${id}`);
    }
  };

  return (
    <div>
      <h1 className="text-xl font-bold mb-4">Setup Menu</h1>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          submit();
        }}
      >
        <input
          className="border px-2 py-1 mr-2"
          value={data.name}
          onChange={(e) => setData("name", e.target.value)}
          placeholder="Nama menu"
        />
        <input
          className="border px-2 py-1 mr-2"
          value={data.uri}
          onChange={(e) => setData("uri", e.target.value)}
          placeholder="/uri"
        />
        <input
          className="border px-2 py-1 mr-2"
          value={data.parent_id ?? ""}
          onChange={(e) =>
            setData(
              "parent_id",
              e.target.value ? parseInt(e.target.value) : null
            )
          }
          placeholder="Parent ID (optional)"
        />
        <button className="bg-blue-500 text-white px-4 py-1">Simpan</button>
      </form>
    </div>
  );
}
