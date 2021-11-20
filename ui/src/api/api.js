const axios = require("axios");

// const baseUrl = "http://192.168.0.20:8080";
export const baseUrl = "http://localhost:8080";
// const baseUrl = "http://localhost:8080";

export const fetchGame = (id) => axios.get(baseUrl + "/game/" + id);
export const newGame = () => axios.get(baseUrl + "/game/new");
export const joinGame = (id) => axios.get(baseUrl + "/game/" + id + "/join");
export const startGame = (id) => axios.get(baseUrl + "/game/" + id + "/start");
export const playerSession = (gid, id) =>
  axios.get(baseUrl + `/game/${gid}/player/${id}`);
export const playerLeave = (gid, id) =>
  axios.get(baseUrl + `/game/${gid}/player/${id}/leave`);
export const playerFold = (gid, id) =>
  axios.get(baseUrl + `/game/${gid}/player/${id}/fold`);
export const proceedWithGame = (id) =>
  axios.get(baseUrl + "/game/" + id + "/proceed");
