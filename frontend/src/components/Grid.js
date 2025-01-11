import { useState } from "react";
import "./Grid.css";

const Cell = ({ value }) => {
  const [place, setPlace] = useState("");
  const [disable, setDisable] = useState(false);
  const onClick = () => {
    setDisable(true);
    setPlace(value);
  };
  return (
    <button onClick={onClick} disabled={disable}>
      {place}
    </button>
  );
};

const Grid = ({ value }) => {
  return (
    <div className="grid-container">
      <div>
        <Cell value={value} />
        <Cell value={value} />
        <Cell value={value} />
      </div>
      <div>
        <Cell value={value} />
        <Cell value={value} />
        <Cell value={value} />
      </div>
      <div>
        <Cell value={value} />
        <Cell value={value} />
        <Cell value={value} />
      </div>
    </div>
  );
};

export default Grid;
