import { useState } from "react";
import Client from "../client/game_client";
import "./Grid.css";

const client = new Client("127.0.0.1", ":3002");

const Cell = ({ value, i, j }) => {
  const [place, setPlace] = useState("");
  const [disable, setDisable] = useState(false);

  const onClick = () => {
    if (value !== undefined || value !== "") {
      setDisable(true);
      setPlace(value);
      client.sendMove(i, j);
    }
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
        <Cell value={value} i={0} j={0} />
        <Cell value={value} i={0} j={1} />
        <Cell value={value} i={0} j={2} />
      </div>
      <div>
        <Cell value={value} i={1} j={0} />
        <Cell value={value} i={1} j={1} />
        <Cell value={value} i={1} j={2} />
      </div>
      <div>
        <Cell value={value} i={2} j={0} />
        <Cell value={value} i={2} j={1} />
        <Cell value={value} i={2} j={2} />
      </div>
    </div>
  );
};

export default Grid;
