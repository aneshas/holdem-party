import { useState } from "react";
import { Button, Container } from "react-bootstrap";
import { useParams } from "react-router";
import { fetchGame, newGame, proceedWithGame, startGame } from "../api/api";
import Card, { DeckCard } from "./Card";
import { useInterval } from "./hooks";
import Player from "./Player";

const PokerTable = () => {
  const [game, setGame] = useState({});
  const { id } = useParams();

  useInterval(() => {
    fetchGame(id).then((resp) => {
      if (!resp.data) {
        return;
      }

      setGame(resp.data);
    });
  }, 200);

  const haveMinimumPlayers = () => game.Players && game.Players.length > 1;
  return (
    <>
      <small className="text-center text-white p-1">
        <b>Join Link: </b> {`http://localhost:3000/table/${id}/join`}
      </small>

      <Container fluid className="mt-5">
        <div
          className="d-flex justify-content-center"
          style={{
            minHeight: "192px",
          }}
        >
          {game.Players?.map((p, i) => {
            return (
              <Player
                hands={game.Hands}
                blinds={game.Blinds}
                showHand={!game.IsStarted}
                key={i}
                id={id}
                pid={p.ID}
                number={i}
              />
            );
          })}
        </div>
      </Container>

      <Container fluid className="d-flex  justify-content-center">
        <div
          className="d-flex my-5 mx-5"
          style={{
            minWidth: "1400px",
          }}
        >
          <DeckCard />
          {game.Flop?.map((c, i) => {
            return <Card key={i} card={c} />;
          })}
          {!!game.Turn && <Card card={game.Turn} />}
          {!!game.River && <Card card={game.River} />}
        </div>
      </Container>

      <Container fluid className="d-flex  align-items-center flex-column">
        {game.IsStarted && <NextButton id={id} />}
        {game.IsStarted && <NewGameButton />}
        {!game.IsStarted && haveMinimumPlayers() && <StartButton id={id} />}
        {!haveMinimumPlayers() && (
          <span className="text-info">
            At least two players needed to start the game
          </span>
        )}
      </Container>
    </>
  );
};

const NewGameButton = () => {
  const handleNew = () => {
    newGame().then((resp) => {
      window.open(`/table/${resp.data.ID}`, "_blank");
    });
  };

  return (
    <Button className="mt-5" onClick={handleNew} variant="danger" size="lg">
      New Game
    </Button>
  );
};

const NextButton = ({ id }) => {
  const handleClick = () => {
    proceedWithGame(id).then(console.log);
  };

  return (
    <Button onClick={handleClick} variant="warning" size="lg">
      Deal / Showdown
    </Button>
  );
};

const StartButton = ({ id }) => {
  const handleClick = () => {
    startGame(id).then(console.log);
  };

  return (
    <Button onClick={handleClick} variant="warning" size="lg">
      Start New Round
    </Button>
  );
};

export default PokerTable;
