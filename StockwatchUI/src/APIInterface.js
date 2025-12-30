import axios from "axios";
const STOCKWATCH_API_URL = import.meta.env.VITE_STOCKWATCH_API_URL

export async function getWatchlists(){
    try {
        const response = await axios.post(
            `${STOCKWATCH_API_URL}/GetWatchlists`,
            { ID: "eb0dcdff-741d-437c-ad64-35b267a91494" },
            { headers: { "Content-Type": "application/json" } }
        );
      
          return response.data; // returns the data as a JS object
      }catch (error) {
          console.error(error);
          return {}; // return empty object on error
      }
}