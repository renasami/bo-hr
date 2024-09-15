import { FC } from "react";
import { Layout } from "../layout/layout";
import { Tabs } from "@mantine/core";

export const Branch: FC = () => {
  return (
    <Layout>
      <h1>branch</h1>
      <Tabs defaultValue="gallery">
        <Tabs.List>
          <Tabs.Tab value="all">all</Tabs.Tab>
          <Tabs.Tab value="create">create</Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="all">Gallery tab content</Tabs.Panel>

        <Tabs.Panel value="create">Messages tab content</Tabs.Panel>
      </Tabs>
    </Layout>
  );
};
