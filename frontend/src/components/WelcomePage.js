import { useState } from "react";
import LoginForm from "./Forms/Login";
import SignUp from "./Forms/SignUp";
import "./WelcomePage.css";
import Grid from "./Grid";
import UserDetails from "./UserDetails";
import { RemoveUser } from "../store/reducers";
import { useDispatch, useSelector } from "react-redux";
const WelcomePage = () => {
  const [loginForm, setLoginForm] = useState(true);
  const [playing, setPlaying] = useState(false);
  const toggle = () => {
    setLoginForm(!loginForm);
  };
  const dispatch = useDispatch();
  const { user, status } = useSelector((state) => state);
  return (
    <div className="welcome">
      <h1>Tic Tac Toe</h1>
      <p>Play with real world players, be the best and climb up the ranks</p>
      <p>you are playing against : </p>
      <Grid value={"O"} />

      <div style={{ textAlign: "right" }}>
        <div>
          {status.isLoggedIn ? (
            <>
              <UserDetails user={user} />
              <button className="toggle-button" onClick={dispatch(RemoveUser)}>
                logout
              </button>
            </>
          ) : (
            <>
              <div style={{ display: "inline-block", textAlign: "left" }}>
                {loginForm ? <LoginForm /> : <SignUp />}
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
