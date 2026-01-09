import React from "react";

import { useState, useEffect} from "react";
import Login from "./components/Login";
import Header from "./components/Header";
import WatchlistGrid from "./components/WatchlistGrid";
import AddWatchlistPopup from "./components/AddWatchlistPopup";
import ErrorPopup from "./components/ErrorMSGPopup";
import WarningPopup from "./components/WarningMSGPopup";
import { checkHealth, login, getWatchlists, deleteWatchlist, createWatchlist } from "./APIInterface"

export default function App() {
    const [serverStatus, setServerStatus] = useState(false);
    const [errorPopup, setErrorPopup] = useState(false);
    const [warningPopup, setWarningPopup] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [warningMessage, setWarningMessage] = useState("");
    const [userEmail, setUserEmail] = useState(null);
    const [userID, setUserId] = useState("");
    const [watchlists, setWatchlists] = useState([]);
    const [showAddWatchlistPopup, setShowAddPopup] = useState(false);
    const [delWID, setDelWID] = useState(null); // wathclist ID when the user wants to delete one

    async function healthCheck(){
        const health = await checkHealth();
        
        if (health === null) {
            setErrorPopup(true)
            setErrorMessage("Server is Down")
            return
        }
        
        console.log("Server Health Check:", health);
        setServerStatus(true)
    }

    async function handleLogin(email) {
        var login_userID = await login(email)
        
        if (login_userID){
            localStorage.setItem("userID", login_userID.userID);
            localStorage.setItem("userEmail", email);

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
        localStorage.removeItem("userEmail");
    }

    async function fetchWatchlists() {
        const data = await getWatchlists(userID);
        setWatchlists(data);
    }

    async function init() {
        await fetchWatchlists();
    }
    
    useEffect(() => {
        healthCheck()
        const savedUserID = localStorage.getItem("userID");
        const savedUserEmail = localStorage.getItem("userEmail");

        if (savedUserID) {
            setUserId(savedUserID);
            setUserEmail(savedUserEmail)
        }
    }, []);

    useEffect(() => {
        if(userID){
            init(); 
        }
    }, [userID]);

    const addNewWatchlist = async (WatchlistName, newAlerts) => {
        if (newAlerts.length == 0){
            setErrorPopup(true)
            setErrorMessage("New watchlists must require 1 alert.")
            return
        }
        
        var newWatchlistDataJson = {
            "userID": userID,
            "email" : userEmail,
            "watchlistData": {
                "name": WatchlistName,
                "alerts": newAlerts
            }
        }
        await createWatchlist(newWatchlistDataJson)
        setShowAddPopup(false)
        await  fetchWatchlists();
    }

    const delWatchlistConfirm = (watchlistID, watchlistName) =>{
        setWarningPopup(true)
        setWarningMessage("Are you sure you want to delete " + watchlistName)
        setDelWID(watchlistID)
    }

    const delWatchlist = async () => {
        if (delWID == "" || delWID === null) return
        
        var deleteWLJson = {
            "ID": delWID
        }
        await deleteWatchlist(deleteWLJson)
        setDelWID(null)
        await fetchWatchlists();
    }

    if (!userID) {
        return <>
        <Login handleLogin={handleLogin} />;
            {errorPopup && (
                <ErrorPopup 
                    message={errorMessage}
                    onClose={() => setErrorPopup(false)} 
                />
            )}
        </>
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
            {warningPopup && (
                <WarningPopup 
                    message={warningMessage}
                    onAccept={() => {delWatchlist(); setWarningPopup(false)}}
                    onCancel={() => setWarningPopup(false)}
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
                onDelete={delWatchlistConfirm}
            />
        </>
    );
}
