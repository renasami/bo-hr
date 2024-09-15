import { FC } from "react";
import { Layout } from "../layout/layout";
import "@mantine/core/styles.css";

export const Users: FC = () => {
  const hoge = "users";

  return (
    <Layout>
      <h1>{hoge}</h1>
    </Layout>
  );
};
