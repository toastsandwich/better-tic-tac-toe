import "./WelcomePage.css";
import axios from "axios";
const FindMatch = ({ user }) => {
  const onClick = () => {
    const email = user.email;
    axios
      .post(`http://localhost:3001/api/findMatch?email=${email}`)
      .then((resp) => {
        if (resp.data) {
          console.log(resp.data);
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <p>
      Click here to find match :
      <button className="toggle-button" onClick={onClick}>
        Find Match
      </button>
    </p>
  );
};

export default FindMatch;
