import "./WelcomePage.css";
import axios from "axios";
const FindMatch = ({ user, setGameSide, token }) => {
  const onClick = () => {
    const email = user.email;
    axios
      .post(
        `http://localhost:3001/api/findMatch`,
        {
          email: email,
        },
        {
          headers: {
            Authorization: token,
          },
        }
      )
      .then((resp) => {
        if (resp.data) {
          console.log(resp.data.face);
          setGameSide(resp.data.face);
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
