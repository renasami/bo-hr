import { customFetch } from "./util";

const base = import.meta.env.VITE_HR_URL;

export const createUserTitle = async (
  tenantNameId: string,
  titleName: string
) => {
  const url = "/user_title";

  const options: RequestInit = {
    method: "POST",
    headers: {
      "content-type": "application/json",
    },
    body: JSON.stringify({ tenantNameId, titleName }),
  };

  return await customFetch(base, url, options);
};
