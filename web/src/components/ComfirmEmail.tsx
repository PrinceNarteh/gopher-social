import { API_URL } from "../App";

export const ConfirmEmail = () => {
  const token = "";
  const handleConfirm = async () => {
    const res = await fetch(`${API_URL}/users/activate/${token}`, {
      method: "PUT",
    });

    if (res.ok) {
      const data = await res.json();
      console.log("Email activated successfully:", data);
      // Redirect or show success message
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
