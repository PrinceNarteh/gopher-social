import { API_URL } from "../App";
import { useNavigate, useParams } from "react-router-dom";

export const ConfirmEmail = () => {
  const { token = "" } = useParams<{ token?: string }>();
  const redirect = useNavigate();

  const handleConfirm = async () => {
    const res = await fetch(`${API_URL}/users/activate/${token}`, {
      method: "PUT",
    });

    if (res.ok) {
      const data = await res.json();
      console.log("Email activated successfully:", data);
      redirect("/"); // Redirect to home or another page after activation
    } else {
      const errorData = await res.json();
      console.error("Error activating email:", errorData);
      // Show error message to the user
    }
  };

  return (
    <div>
      <h2>Activate Email</h2>
      <button onClick={handleConfirm}>Activate</button>
    </div>
  );
};
