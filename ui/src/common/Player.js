import { useState } from "react";
import { playerLeave, playerSession } from "../api/api";
import back from "../assets/img/1B.svg";
import backRed from "../assets/img/2B.svg";
import Card from "./Card";
import { useInterval } from "./hooks";
import "./Player.css";
import { Button } from "react-bootstrap";

const Player = ({ id, number, hands, blinds, showHand = false }) => {
  const hand = hands ? hands[id] : null;
  const isSmall = blinds && blinds[0] === id;
  const isBig = blinds && blinds[1] === id;

  const [holdsHand, setHoldsHand] = useState(false);

  useInterval(() => {
    playerSession(id).then((resp) => {
      if (resp.data.Hand) {
        setHoldsHand(true);
      } else {
        setHoldsHand(false);
      }
    });
  }, 1000);

  const variant = holdsHand ? "text-warning" : "text-danger";
  const img = holdsHand ? back : backRed;

  const handleLeave = () => {
    playerLeave(id).then(console.log);
  };

  return (
    <div className="player px-3">
      <div
        className={`text-uppercase ${variant} d-flex justify-content-between align-items-center`}
      >
        <b>P{number + 1}</b>
        <Button onClick={handleLeave} variant="link" className="text-danger">
          Remove
        </Button>
      </div>
      {showHand && hand ? (
        <div className="d-flex hand">
          <Card card={hand[0]} />
          <Card card={hand[1]} />
        </div>
      ) : (
        <div>
          <img src={img} alt="card" />
          <img src={img} alt="card" />
        </div>
      )}

      {isSmall && <b className="text-white">SMALL</b>}
      {isBig && <b className="text-info">BIG</b>}
    </div>
  );
};

export default Player;
