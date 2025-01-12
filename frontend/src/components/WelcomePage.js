import { useEffect, useState } from "react";
import LoginForm from "./Forms/Login";
import SignUp from "./Forms/SignUp";
import "./WelcomePage.css";
import Grid from "./Grid";
import UserDetails from "./UserDetails";
import { useDispatch, useSelector } from "react-redux";
import { RemoveUser } from "../store/reducers";

const WelcomePage = () => {
  const [loginForm, setLoginForm] = useState(true);
  const [authToken, setAuthToken] = useState(null); // Changed to null for initial state
  const dispatch = useDispatch();
  const { user } = useSelector((state) => state); // Adjust based on your state structure

  useEffect(() => {
    const token = localStorage.getItem("authtoken");
    setAuthToken(token); // This will set authToken immediately upon component mount
  }, [authToken]);

  const toggle = () => setLoginForm(!loginForm);

  const handleLogout = () => {
    dispatch(RemoveUser());
    setAuthToken(null);
    localStorage.removeItem("authtoken");
  };

  // Check for authToken to determine logged-in state
  const isLoggedIn = authToken !== null;

  return (
    <div className="welcome">
      <h1>Tic Tac Toe</h1>
      <p>Play with real world players, be the best and climb up the ranks</p>
      <p>you are playing against : </p>
      <Grid value={"O"} />

      <div style={{ textAlign: "right" }}>
        <div>
          {isLoggedIn ? (
            <>
              <UserDetails user={user} />
              <button className="toggle-button" onClick={handleLogout}>
                Logout
              </button>
            </>
          ) : (
            <>
              <div style={{ display: "inline-block", textAlign: "left" }}>
                {loginForm ? <LoginForm /> : <SignUp />}
              </div>
              <br />
              <button className="toggle-button" onClick={toggle}>
                {loginForm ? "Create an account?" : "Already have an account?"}
              </button>
            </>
          )}
        </div>
      </div>
    </div>
  );
};

export default WelcomePage;
