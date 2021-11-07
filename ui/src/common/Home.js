import { Container, Button } from "react-bootstrap";
import { newGame } from "../api/api";

const Home = () => {
  const handleNew = () => {
    newGame().then((resp) => {
      window.location.href = `/table/${resp.data.ID}`;
    });
  };

  return (
    <Container fluid className="text-center pt-5">
      <h1 className="text-white">Texas HoldEm Poker</h1>
      <Button onClick={handleNew} size="lg" variant="warning" className="my-3">
        New Game
      </Button>
    </Container>
  );
};

export default Home;
