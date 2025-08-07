import "./App.css";
import { ConfirmEmail } from "./components/ComfirmEmail";

export const API_URL =
  import.meta.env.VITE_API_URL || "http://localhost:4000/v1";

function App() {
  return <ConfirmEmail />;
}

export default App;
