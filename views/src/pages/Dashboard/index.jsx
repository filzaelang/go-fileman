import Layout from "../../layouts/Layout";

function Index({ phrase }) {
  return (
    <>
      <h1 className="font-bold text-2xl">{phrase}</h1>
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
