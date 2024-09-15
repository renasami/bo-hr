import { FC } from "react";
import { Layout } from "../layout/layout";
import "@mantine/core/styles.css";

export const CreatePage: FC = () => {
  const hoge = "hoge";
  return (
    <Layout>
      <h1>{hoge}</h1>
    </Layout>
  );
};
