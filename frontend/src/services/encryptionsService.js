import axios from 'axios'
const baseUrl = '/api/protected'

let token = null

const setToken = newToken => {
  token = `bearer ${newToken}`
}

const getCaesars = async () => {
    const config = {
        headers: { Authorization: token }
      }
      const response = await axios.get(`${baseUrl}/caesar`, config)
      return response
}

const encryptionService = { setToken, getCaesars }

export default encryptionService