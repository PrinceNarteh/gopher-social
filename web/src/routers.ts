import { createBrowserRouter } from "react-router-dom";
import App from "./App.tsx";
import { ConfirmEmail } from "./components/ComfirmEmail.tsx";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/confirm/:token",
    element: <ConfirmEmail />,
  },
]);
