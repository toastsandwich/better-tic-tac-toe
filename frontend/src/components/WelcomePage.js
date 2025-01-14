import { useEffect, useState } from "react";
import LoginForm from "./Forms/Login";
import SignUp from "./Forms/SignUp";
import "./WelcomePage.css";
import Grid from "./Grid";
import UserDetails from "./UserDetails";
import { useDispatch, useSelector } from "react-redux";
import { RemoveUser } from "../store/reducers";
import FindMatch from "./FindMatch.js";
const WelcomePage = () => {
  const [loginForm, setLoginForm] = useState(true);
  const [authToken, setAuthToken] = useState(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const dispatch = useDispatch();
  const { user } = useSelector((state) => state);

  useEffect(() => {
    const token = localStorage.getItem("authtoken");
    setAuthToken(token);
  }, [authToken]);

  const toggle = () => setLoginForm(!loginForm);

  const handleLogout = () => {
    setIsLoggedIn(false);
    dispatch(RemoveUser());
    setAuthToken(null);
    localStorage.removeItem("authtoken");
  };

  return (
    <div className="welcome">
      <h1>Tic Tac Toe</h1>
      <p>Play with real world players, be the best and climb up the ranks</p>
      <FindMatch user={user} />
      <Grid value={"O"} />

      <div style={{ textAlign: "right" }}>
        <div>
          {authToken ? (
            <>
              <UserDetails user={user} />
              <button className="toggle-button" onClick={handleLogout}>
                Logout
              </button>
            </>
          ) : (
            <>
              <div style={{ display: "inline-block", textAlign: "left" }}>
                {loginForm ? (
                  <LoginForm
                    isLoggedIn={isLoggedIn}
                    setIsLoggedIn={setIsLoggedIn}
                  />
                ) : (
                  <SignUp />
                )}
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
