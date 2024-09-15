import { FC } from "react";
import { Layout } from "../layout/layout";
import "@mantine/core/styles.css";
import { Button, Input, Tabs, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { createUserTitle } from "../api/hr";
import "@mantine/notifications/styles.css";
import { notifications } from "@mantine/notifications";
export const UserTitle: FC = () => {
  const form = useForm({
    initialValues: {
      tenantNameId: "",
      title: "",
    },

    validate: {
      // title: (value) =>
      //   value.length > 20 && value.length === 0 ? null : "invalid title name",
      // tenantNameId: (value) =>
      //   value.length > 20 && value.length === 0
      //     ? null
      //     : "invalid tenant name id",
    },
  });

  const handleSubmit = async (value: {
    tenantNameId: string;
    title: string;
  }) => {
    try {
      await createUserTitle(value.tenantNameId, value.title);
    } catch {
      notifications.show({
        title: "保存に失敗しました",
        message: "保存失敗",
      });
      return;
    }
    notifications.show({
      message: "成功！",
    });
  };

  return (
    <Layout>
      <h1>役職</h1>
      <Tabs defaultValue="gallery">
        <Tabs.List>
          <Tabs.Tab value="all">all</Tabs.Tab>
          <Tabs.Tab value="create">create</Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="all">Gallery tab content</Tabs.Panel>

        <Tabs.Panel value="create">
          <form
            onSubmit={form.onSubmit((value) => {
              console.log(value);
              handleSubmit(value);
            })}
          >
            <TextInput
              withAsterisk
              label="テナント名"
              placeholder="yonyon"
              key={form.key("tenantNameId")}
              {...form.getInputProps("tenantNameId")}
            />
            <TextInput
              withAsterisk
              label="役職"
              placeholder="部長"
              key={form.key("title")}
              {...form.getInputProps("title")}
            />
            <br />
            <Button type="submit">作成</Button>
          </form>
        </Tabs.Panel>
      </Tabs>
    </Layout>
  );
};
