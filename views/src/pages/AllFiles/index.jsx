import Layout from "../../layouts/Layout";

function index({ phrase }) {
  return (
    <>
      <h1 className="font-bold text-2xl">{phrase}</h1>
    </>
  );
}

index.layout = (page) => <Layout>{page}</Layout>;
export default index;
