import axios from "axios";
import "./styles.css";
import { useState } from "react";

const SignUp = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [country, setCountry] = useState("");
  const [passsord, setPassword] = useState("");

  const submit = () => {
    axios
    .post("http://localhost:3001/api/user/create", {
        country: country,
        email: email,
        username: username,
        password: passsord,
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <form className="signup-form">
      <h3 className="signup-title">Sign Up</h3>
      <label htmlFor="country" className="label-countery">
        Country
      </label>
      <input
        id="country"
        type="text"
        name="country"
        className="input-country"
        onChange={(e) => {
          setCountry(e.target.value);
        }}
        required
      />
      <label htmlFor="email" className="label-email">
        Email
      </label>
      <input
        id="email"
        type="email"
        name="email"
        className="input-email"
        onChange={(e) => {
          setEmail(e.target.value);
        }}
        required
      />
      <label htmlFor="username" className="label-username">
        Username
      </label>
      <input
        id="username"
        type="text"
        name="username"
        className="input-username"
        onChange={(e) => {
          setUsername(e.target.value);
        }}
        required
      />
      <label htmlFor="password" className="label-password">
        Password
      </label>
      <input
        id="password"
        type="password"
        name="password"
        className="input-password"
        onChange={(e) => {
          setPassword(e.target.value);
        }}
        required
      />
      <button type="button" className="btn-signup" onClick={submit}>
        signup
      </button>
    </form>
  );
};

export default SignUp;
