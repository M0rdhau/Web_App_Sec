
import axios from 'axios'
const baseUrl = '/api/public/login'

const login = async (username, password) => {
  const response = await axios.post(baseUrl, { username, password })
  return response.data
}

export default login
