import axios from "axios";
import { useState } from "react";
import Swal from "sweetalert2";
import { SetUser } from "../../store/reducers";
import { useDispatch } from "react-redux";

const Login = ({ isLoggedIn, setIsLoggedIn }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const dispatch = useDispatch();
  const submit = () => {
    if (!email || !password) {
      Swal.fire({
        position: "top-right",
        title: "Please fill in all fields",
        showConfirmButton: false,
        timer: 500,
        width: "80%",
        customClass: {
          popup: "custom-height-modal-fail",
        },
      });
      return;
    }
    axios
      .post("http://localhost:3001/api/login", {
        email: email,
        password: password,
      })
      .then((resp) => {
        const { message, token, user } = resp.data;
        if (message === "success") {
          dispatch(SetUser({ ...user }));
          localStorage.setItem("authtoken", token);
          setIsLoggedIn(true);
        }
        Swal.fire({
          position: "top-right",
          title: "login success",
          showConfirmButton: false,
          timer: 500,
          width: "80%",
          customClass: {
            popup: "custom-height-modal-success",
          },
        });
      })
      .catch(() => {
        Swal.fire({
          position: "top-right",
          title: "login fail",
          showConfirmButton: false,
          timer: 500,
          width: "80%",
          customClass: {
            popup: "custom-height-modal-fail",
          },
        });
      });
  };

  return (
    <form className="auth-form">
      <h3 className="auth-title">Login</h3>
      <label htmlFor="email" className="label-email">
        Email
      </label>
      <input
        id="email"
        type="email"
        name="email"
        className="input-field"
        onChange={(e) => setEmail(e.target.value)}
      />
      <label htmlFor="password" className="label-password">
        Password
      </label>
      <input
        id="password"
        type="password"
        name="password"
        className="input-field"
        onChange={(e) => setPassword(e.target.value)}
      />
      <button type="button" className="btn-auth" onClick={submit}>
        Login
      </button>
    </form>
  );
};

export default Login;
