import axios from "axios";
const STOCKWATCH_API_URL = import.meta.env.VITE_STOCKWATCH_API_URL

export async function checkHealth() {
  try {
    const response = await fetch(`${STOCKWATCH_API_URL}/Health`);

    if (!response.ok) {
      throw new Error("Server returned an error");
    }

    const text = await response.text(); // "ok"
    return text;
  } catch (err) {
    console.error("Health check failed:", err);
    return null;
  }
}

export async function login(email) {
    try {
        const response = await axios.post(
            `${STOCKWATCH_API_URL}/LoginRequest`,
            { email: email },
            { headers: { "Content-Type": "application/json" } }
        );
        return response.data;
    } catch (err) {
        console.error("Health check failed:", err);
        return null;
    }
}


export async function getWatchlists(userID){
    try {
        const response = await axios.post(
            `${STOCKWATCH_API_URL}/GetWatchlists`,
            { ID: userID },
            { headers: { "Content-Type": "application/json" } }
        );
    
        return response.data;
    }catch (error) {
        console.error(error);
        return {};
    }
}

export async function createWatchlist(watchlistData) {
    console.log(watchlistData)
    try {
        const response = await axios.post(
            `${STOCKWATCH_API_URL}/CreateWatchlist`,
            watchlistData,
            { headers: { "Content-Type": "application/json" } }
        );

        return response.data;
    } catch (error) {
        console.error(error);
        return {};
    }
}

export async function deleteWatchlist(watchlistData){
    try {
        const response = await axios.post(
            `${STOCKWATCH_API_URL}/DeleteWatchlist`,
            watchlistData,
            { headers: { "Content-Type": "application/json" } }
        );
    
        return response.data; // returns the data as a JS object
    }catch (error) {
        console.error(error);
        return {}; // return empty object on error
    }
}