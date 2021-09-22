import axios from 'axios';

const host = `http://localhost:3001/api`;
const apiClient = axios.create({
  baseURL: host,
  withCredentials: true,
});

export default apiClient;
