import bside from "../assets/img/1B.svg";

export const DeckCard = () => {
  return (
    <div className="px-1">
      <img src={bside} alt="card" />
    </div>
  );
};

const Card = ({ card }) => {
  let cardName = `${card.Rank}${card.Suit[0]}`;

  cardName = cardName.replace("12", "J");
  cardName = cardName.replace("13", "Q");
  cardName = cardName.replace("14", "K");

  const cardImg = `/assets/img/${cardName}.svg`;

  return (
    <div className="px-1">
      <img src={cardImg} alt="card" />
    </div>
  );
};

export default Card;
