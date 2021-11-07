import { useEffect } from "react";
import { useParams } from "react-router";
import { joinGame } from "../api/api";

const JoinGame = () => {
  const { id } = useParams();

  useEffect(() => {
    joinGame(id).then((resp) => {
      if (!resp) return;

      window.location.href = `/table/${id}/player/${resp.data.PlayerID}`;
    });
  }, [id]);

  return <></>;
};

export default JoinGame;
