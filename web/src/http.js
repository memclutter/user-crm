import Axios from "axios";

const http = Axios.create({
  baseURL: "/api/",
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  }
});

export default http;