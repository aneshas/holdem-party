import { useState } from "react";
import { Button, Container } from "react-bootstrap";
import { useParams } from "react-router";
import { playerFold, playerSession } from "../api/api";
import Card from "./Card";
import { useInterval } from "./hooks";
import "./Seat.css";
import img from "../assets/img/1B.svg";

const Seat = () => {
  const [player, setPlayer] = useState(null);
  const { id, pid } = useParams();

  useInterval(() => {
    playerSession(id, pid).then((resp) => {
      if (!resp.data?.SeatNumber) {
        window.location.href = `/`;
        return;
      }

      setPlayer(resp.data);
    });
  }, 1000);

  const handleFold = () => {
    playerFold(id, pid).then(console.log);
  };

  return (
    <Container className="seat pt-5 d-flex align-items-center flex-column">
      <h1 className="text-warning text-uppercase">
        Player {player?.SeatNumber}
      </h1>

      {player && player.Hand && <Hand hand={player.Hand} />}

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

const Hand = ({ hand }) => {
  const [show, setShow] = useState(false);

  return (
    <div onTouchStart={() => setShow(true)} onTouchEnd={() => setShow(false)}>
      {show && (
        <div className="hand d-flex justify-content-center mt-5">
          <Card card={hand[0]} />
          <Card card={hand[1]} />
        </div>
      )}

      {!show && (
        <div className="hand d-flex justify-content-center mt-5">
          <img src={img} alt="card" />
          <img src={img} alt="card" />
        </div>
      )}
    </div>
  );
};

export default Seat;
