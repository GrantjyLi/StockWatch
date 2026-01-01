import React from "react";

import { useState, useEffect} from "react";
import Header from "./components/Header";
import WatchlistGrid from "./components/WatchlistGrid";
import AddWatchlistPopup from "./components/AddWatchlistPopup";
import ErrorPopup from "./components/ErrorMSGPopup";
import { checkHealth, getWatchlists, deleteWatchlist, createWatchlist } from "./APIInterface"

export default function App() {
    const [serverStatus, setServerStatus] = useState(false);
    const [errorPopup, setErrorPopup] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [userID, setUserId] = useState("eb0dcdff-741d-437c-ad64-35b267a91494");
    const [watchlists, setWatchlists] = useState([]);
    const [showAddWatchlistPopup, setShowAddPopup] = useState(false);

    async function fetchWatchlists() {
        const data = await getWatchlists(userID); // assume this returns JSON
        setWatchlists(data);
    }

    async function init() {
        const health = await checkHealth();
        
        if (health === null) {
            setErrorPopup(true)
            setErrorMessage("Server is Down")
            return; // stop here
        }
        
        console.log("Server Health Check:", health);
        setServerStatus(true)
        await fetchWatchlists();
    }

    useEffect(() => {
        init();
    }, []); // empty array → runs only once on page refresh

    const addNewWatchlist = (WatchlistName, newAlerts) => {
        var newWatchlistDataJson = {
            "userID": userID,
            "watchlistData": {
                "name": WatchlistName,
                "alerts": newAlerts
            }
        }
        createWatchlist(newWatchlistDataJson)
        setShowAddPopup(false)
    }

    const delWatchlist = (watchlistID) => {
        var deleteWLJson = {
            "ID": watchlistID
        }
        deleteWatchlist(deleteWLJson)
    }

    return (
      <>
          <Header />
            {errorPopup && (
                <ErrorPopup 
                    message={errorMessage}
                    onClose={() => setErrorPopup(false)} 
                />
            )}
            {showAddWatchlistPopup && (
                <AddWatchlistPopup 
                    addNewWatchlist={addNewWatchlist}
                    onClose={() => setShowAddPopup(false)} 
                />
            )}
            <WatchlistGrid
                watchlists={watchlists}
                onAdd={()=>setShowAddPopup(true)}
                onDelete={delWatchlist}
            />
        </>
    );
}
