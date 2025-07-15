import axios from 'axios';

const API_URL = 'http://localhost:4000/api/v1'; // Replace with your backend base URL

export const fetchFarms = async () => {
  const response = await axios.get(`${API_URL}/farms`);
   return response.data.farms; 
};


 