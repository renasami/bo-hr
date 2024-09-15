import "@mantine/core/styles.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { UserTitle } from "./pages/CreatePage";
import { Route } from "./pages/Route";
import { Branch } from "./pages/Branch";
import { Users } from "./pages/Users";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <UserTitle />,
    },
    {
      path: "/user-title",
      element: <UserTitle />,
    },
    {
      path: "/branch",
      element: <Branch />,
    },
    {
      path: "/route",
      element: <Route />,
    },
    {
      path: "/users",
      element: <Users />,
    },
  ]);

  return <RouterProvider router={router} />;
}

export default App;
