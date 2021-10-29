const axios = require("axios");

// const baseUrl = "http://192.168.0.20:8080";
const baseUrl = "http://localhost:8080";

export const fetchGame = () => axios.get(baseUrl + "/game");
export const newGame = () => axios.get(baseUrl + "/game/new");
export const joinGame = () => axios.get(baseUrl + "/game/join");
export const startGame = () => axios.get(baseUrl + "/game/start");
export const playerSession = (id) => axios.get(baseUrl + `/game/player/${id}`);
export const playerLeave = (id) =>
  axios.get(baseUrl + `/game/player/${id}/leave`);
export const playerFold = (id) =>
  axios.get(baseUrl + `/game/player/${id}/fold`);
export const proceedWithGame = () => axios.get(baseUrl + "/game/proceed");
