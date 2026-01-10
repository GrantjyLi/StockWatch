import React from "react";

import { useState, useEffect} from "react";
import Login from "./components/Login";
import Header from "./components/Header";
import WatchlistGrid from "./components/WatchlistGrid";
import AddWatchlistPopup from "./components/AddWatchlistPopup";
import ErrorPopup from "./components/ErrorMSGPopup";
import WarningPopup from "./components/WarningMSGPopup";
import { checkHealth, login, createUser, getWatchlists, deleteWatchlist, createWatchlist } from "./APIInterface"

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

    function handleErrorPopup(errorMessage){
        setErrorPopup(true)
        setErrorMessage(errorMessage)
    }

    async function healthCheck(){
        const health = await checkHealth();
        
        if (health === null) {
            handleErrorPopup("Server is Down")
            return
        }
        
        console.log("Server Health Check:", health);
        setServerStatus(true)
    }

    async function handleLogin(email) {
        var loginData = {
            "email": email
        }

        var result = await login(loginData)

        if(!result.ok){
            handleErrorPopup("Failed login for email: " + email)
            return
        }
        const login_userID = result.data.userID
        localStorage.setItem("userID", login_userID);
        localStorage.setItem("userEmail", email);

        setUserEmail(email)
        setUserId(login_userID)
    };

    async function handleCreateUser(email) {
        var userData = {
            "email": email
        }
        var result = await createUser(userData)
        
        if (!result.ok){
            handleErrorPopup("Failed to create new user: " + result.data)
            return false
        }
        alert("User with email " + email + " has been created, proceed to login with it.")
        return true
    };

    const handleLogout = () => {
        setUserEmail(null)
        setUserId("")
        localStorage.removeItem("userID");
        localStorage.removeItem("userEmail");
    }

    async function fetchWatchlists() {
        var fetchWLData = {
            "ID": userID
        }
        const result = await getWatchlists(fetchWLData);
        var watchlistData = result.data
        setWatchlists(watchlistData);
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
        var result = await createWatchlist(newWatchlistDataJson)
        setShowAddPopup(false)

        if (!result.ok){
            handleErrorPopup("Could not create watchlist: " + result.data)
            return
        }

        await fetchWatchlists();
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
        var result = await deleteWatchlist(deleteWLJson)
        setDelWID(null)
        if (!result.ok){
            handleErrorPopup("Could not delete watchlist: " + result.data)
            return
        }
        await fetchWatchlists();
    }

    if (!userID) {
        return <>
            <Login 
                handleLogin={handleLogin}
                handleCreateUser={handleCreateUser}
            />;
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
