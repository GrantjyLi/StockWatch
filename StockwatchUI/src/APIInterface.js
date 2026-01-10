import axios from "axios";
const STOCKWATCH_API_URL = import.meta.env.VITE_STOCKWATCH_API_URL

function handleError(err){
    var errData = err.response?.data

    if (errData == undefined || errData === null) return
    
    if (errData.startsWith("ERROR: duplicate key value violates unique constraint \"users_email_key\"")){
        return "Another account with this email address already exists."
    }else{
        return errData
    }
}

export async function checkHealth() {
  try {
    const response = await fetch(`${STOCKWATCH_API_URL}/Health`);

    if (!response.ok) {
      throw new Error("Server returned an error");
    }

    const text = await response.text(); // "ok"
    return text;
  } catch (err) {
    return null;
  }
}

async function apiCall(URL_PATH, data){
    try {
        const response = await axios.post(
            `${STOCKWATCH_API_URL}/${URL_PATH}`,
            data,
            { headers: { "Content-Type": "application/json" } }
        );
        return {
            "ok": true,
            "data": response.data
        }
    } catch (err) {
        return {
            "ok": false,
            "data": handleError(err)
        }
    }
}

export async function login(loginData) {
    var result = await apiCall("LoginRequest", loginData)
    return result
}

export async function createUser(userData) {
    var result = await apiCall("CreateUser", userData)
    return result
}

export async function getWatchlists(fetchWLData) {
    var result = await apiCall("GetWatchlists", fetchWLData)
    return result
}

export async function createWatchlist(watchlistData) {
    var result = await apiCall("CreateWatchlist", watchlistData)
    return result
}

export async function deleteWatchlist(watchlistData) {
    var result = await apiCall("DeleteWatchlist", watchlistData)
    return result
}