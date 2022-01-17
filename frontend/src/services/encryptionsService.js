import axios from "axios";
const baseUrl = "/api/protected";

let token = null;

const setToken = (newToken) => {
    token = `Bearer ${newToken}`;
};

const getEntries = async (entryRoute) => {
    const config = {
        headers: { Authorization: token },
    };
    const response = await axios.get(`${baseUrl}${entryRoute}`, config);
    return response;
};

const addEntry = async (entryData, entryRoute) => {
    const config = {
        headers: { Authorization: token },
    };
    const response = await axios.post(`${baseUrl}${entryRoute}`, entryData, config);
    return response;
};

const deleteEntry = async (entryID, entryRoute) => {
    const config = {
        headers: { Authorization: token },
    };
    const response = await axios.delete(`${baseUrl}${entryRoute}/${entryID}`, config);
    return response;
}

const encryptionService = { setToken, getEntries, addEntry, deleteEntry };

export default encryptionService;
