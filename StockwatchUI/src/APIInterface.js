import axios from "axios";

export async function getWatchlists(){

try {
    const response = await axios.post(
      "http://localhost:8080/GetWatchlists",
      { ID: "eb0dcdff-741d-437c-ad64-35b267a91494" },
      {
        headers: { "Content-Type": "application/json" }
      }
    );
    
    return response.data; // returns the data as a JS object
  } catch (error) {
    console.error(error);
    return {}; // return empty object on error
  }
}