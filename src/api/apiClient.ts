import axios from 'axios';

const host = `${process.env.REACT_APP_API_URI}/api/v1`;
const apiClient = axios.create({
  baseURL: host,
  withCredentials: true,
});

export default apiClient;
