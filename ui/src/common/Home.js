import { Container, Button } from "react-bootstrap";
import { joinGame } from "../api/api";

const Home = () => {
  const handleJoin = () => {
    joinGame().then((resp) => {
      window.location.href = `/player/${resp.data.PlayerID}/${resp.data.PlayerNumber}`;
    });
  };

  return (
    <Container fluid className="text-center pt-5">
      <h1 className="text-white">Texas HoldEm Poker</h1>
      <Button onClick={handleJoin} size="lg" variant="warning" className="my-3">
        Click to Join
      </Button>
    </Container>
  );
};

export default Home;
