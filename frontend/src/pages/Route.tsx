import { FC } from "react";
import { Layout } from "../layout/layout";
import "@mantine/core/styles.css";

export const Route: FC = () => {
  const hoge = "route";
  return (
    <Layout>
      <h1>{hoge}</h1>
    </Layout>
  );
};
