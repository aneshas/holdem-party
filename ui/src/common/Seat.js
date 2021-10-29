import { useState } from "react";
import { Button, Container } from "react-bootstrap";
import { useParams } from "react-router";
import { playerFold, playerSession } from "../api/api";
import Card from "./Card";
import { useInterval } from "./hooks";
import "./Seat.css";

const Seat = () => {
  const [player, setPlayer] = useState(null);
  const { id, seat } = useParams();

  useInterval(() => {
    playerSession(id).then((resp) => {
      setPlayer(resp.data);
    });
  }, 1000);

  const handleFold = () => {
    playerFold(id).then(console.log);
  };

  return (
    <Container className="seat pt-5 d-flex align-items-center flex-column">
      <h1 className="text-warning text-uppercase">Player {seat}</h1>
      {player && player.Hand && (
        <div className="hand d-flex justify-content-center mt-5">
          <Card card={player.Hand[0]} />
          <Card card={player.Hand[1]} />
        </div>
      )}

      {player && player.Hand && (
        <Button
          onClick={handleFold}
          size="lg"
          variant="warning"
          className="mt-5"
        >
          Fold
        </Button>
      )}

      {player && !player.Hand && (
        <span className="text-info">Please wait for the next round</span>
      )}
    </Container>
  );
};

export default Seat;
