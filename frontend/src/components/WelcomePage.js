import { useState } from "react";
import LoginForm from "./Forms/Login";
import SignUp from "./Forms/SignUp";
import "./WelcomePage.css";
import Grid from "./Grid";
import UserDetails from "./UserDetails";
import axios from "axios";
const WelcomePage = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [loginForm, setLoginForm] = useState(true);
  const [playing, setPlaying] = useState(false); //why?
  const [user, setUser] = useState(null);
  const [token, setToken] = useState("");
  const toggle = () => {
    setLoginForm(!loginForm);
  };

  return (
    <div className="welcome">
      <h1>Tic Tac Toe</h1>
      <p>Play with real world players, be the best and climb up the ranks</p>
      <p>you are playing against : </p>
      <Grid value={"O"} />

      <div style={{ textAlign: "right" }}>
        {/* login/signup */}
        <div>
          {isLoggedIn ? (
            <UserDetails user={user} />
          ) : (
            <>
              <div style={{ display: "inline-block", textAlign: "left" }}>
                {loginForm ? (
                  <LoginForm
                    setToken={setToken}
                    setIsLoggedIn={setIsLoggedIn}
                    setUser={setUser}
                  />
                ) : (
                  <SignUp />
                )}
              </div>
              <br />
              <button className="toggle-button" onClick={toggle}>
                {loginForm
                  ? "create an account ?"
                  : "already have an account ?"}
              </button>
            </>
          )}
        </div>
      </div>
    </div>
  );
};

export default WelcomePage;
