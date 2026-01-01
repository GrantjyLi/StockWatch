import React from "react";

import { useState, useEffect} from "react";
import Header from "./components/Header";
import WatchlistGrid from "./components/WatchlistGrid";
import AddWatchlistPopup from "./components/AddWatchlistPopup";
import ErrorPopup from "./components/ErrorMSGPopup";
import { checkHealth, getWatchlists, createWatchlist } from "./APIInterface"

export default function App() {
    const [serverStatus, setServerStatus] = useState(false);
    const [errorPopup, setErrorPopup] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [watchlists, setWatchlists] = useState([]);
    const [showAddWatchlistPopup, setShowAddPopup] = useState(false);

    async function fetchWatchlists() {
        const data = await getWatchlists(); // assume this returns JSON
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
        var newWatchlistData = {
            "name": WatchlistName,
            "alerts": newAlerts
        }

        createWatchlist(newWatchlistData)
        setShowAddPopup(false)
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
            />
        </>
    );
}
