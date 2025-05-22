import Layout from "../../layouts/Layout";
import { usePage } from "@inertiajs/react";

function Index({ phrase }) {
  const { url } = usePage();

  return (
    <>
      <div>
        <h1 className="text-black text-2xl">{url}</h1>
      </div>
      <div className="mt-5">
        <p>GeneralPage</p>
        <p>{phrase}</p>
      </div>
    </>
  );
}

Index.layout = (page) => {
  const props = page.props;
  return <Layout menus={props.menus}>{page}</Layout>;
};

export default Index;
