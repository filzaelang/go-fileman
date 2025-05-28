import Layout from "../../layouts/Layout";
import GeneralPage from "../GeneralPage/Index";

function Index({ phrase, items }) {
  return (
    <>
      {/* <h1 className="font-bold text-2xl">{phrase}</h1> */}
      <GeneralPage phrase={phrase} items={items} />
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
