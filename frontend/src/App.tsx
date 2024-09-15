import "@mantine/core/styles.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { CreatePage } from "./pages/CreatePage";
import { Route } from "./pages/Route";
import { Branch } from "./pages/Branch";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <CreatePage />,
    },
    {
      path: "/user-title",
      element: <CreatePage />,
    },
    {
      path: "/branch",
      element: <Branch />,
    },
    {
      path: "/route",
      element: <Route />,
    },
  ]);

  return <RouterProvider router={router} />;
}

export default App;
