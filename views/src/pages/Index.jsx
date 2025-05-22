/* eslint-disable no-unused-vars */
import { useState } from "react";
import { Link } from "@inertiajs/react";
import Layout from "../layouts/Layout";

function Index({ phrase }) {
  const [count, setCount] = useState(0);

  return (
    <>
      <h1>{phrase}</h1>
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
