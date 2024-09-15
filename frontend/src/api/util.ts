export const customFetch = async <T extends never>(
  base: string,
  path: string,
  options: RequestInit
): Promise<T> => {
  if (base === "" || base === null || base === undefined) {
    throw new Error("some fucking error is occurred");
  }

  const response = await fetch(base + path, options);

  if (!response.ok) {
    throw new Error("cannot fetch");
  }

  // レスポンスをJSONとして解析し、T型の結果を返す
  return (await response.json()) as T;
};
