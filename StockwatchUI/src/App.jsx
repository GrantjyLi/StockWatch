import React from "react";

import { useState, useEffect} from "react";
import Login from "./components/Login";
import Header from "./components/Header";
import WatchlistGrid from "./components/WatchlistGrid";
import AddWatchlistPopup from "./components/AddWatchlistPopup";
import ErrorPopup from "./components/ErrorMSGPopup";
import { checkHealth, login, getWatchlists, deleteWatchlist, createWatchlist } from "./APIInterface"

export default function App() {
    const [serverStatus, setServerStatus] = useState(false);
    const [errorPopup, setErrorPopup] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [userEmail, setUserEmail] = useState(null);
    const [userID, setUserId] = useState("");
    const [watchlists, setWatchlists] = useState([]);
    const [showAddWatchlistPopup, setShowAddPopup] = useState(false);

    async function handleLogin(email) {
        var login_userID = await login(email)
        
        if (login_userID){
            localStorage.setItem("userID", login_userID.userID);

            setUserEmail(email)
            setUserId(login_userID.userID)
        }else{
            alert("Failed login for email: " + email)
        }
    };

    const handleLogout = () => {
        setUserEmail(null)
        setUserId("")
        localStorage.removeItem("userID");

    }

    async function fetchWatchlists() {
        const data = await getWatchlists(userID); // assume this returns JSON
        setWatchlists(data);
    }

    async function init() {
        const health = await checkHealth();
        
        if (health === null) {
            setErrorPopup(true)
            setErrorMessage("Server is Down")
            return
        }
        
        console.log("Server Health Check:", health);
        setServerStatus(true)
        await fetchWatchlists();
    }
    
    useEffect(() => {
        const savedUserID = localStorage.getItem("userID");

        if (savedUserID) {
            setUserId(savedUserID);
        }
    }, []);

    useEffect(() => {
        if(userID){
            init(); 
        }
    }, [userID]); // empty array → runs only once on page refresh

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

    if (!userID) {
        return <Login handleLogin={handleLogin} />;
    }

    return (
        <>
            <Header 
                handleLogout = {handleLogout}
            />
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
