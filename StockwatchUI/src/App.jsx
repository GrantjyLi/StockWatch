import React from "react";

import { useState, useEffect} from "react";
import Header from "./components/Header";
import WatchlistGrid from "./components/WatchlistGrid";
import AddWatchlistPopup from "./components/AddWatchlistPopup";
import { getWatchlists } from "./APIInterface"

export default function App() {
  const [watchlists, setWatchlists] = useState([]);
  const [showAddWatchlistPopup, setShowAddPopup] = useState(false);

  useEffect(() => {
    async function fetchWatchlists() {
      const data = await getWatchlists(); // assume this returns JSON
      setWatchlists(data);
    }

    fetchWatchlists();
  }, []); // empty array → runs only once on page refresh

  const addWatchlist = () => {
    setShowAddPopup(true)
    // setWatchlists([
    //   ...watchlists,
    //   {
    //     id: Date.now(),
    //     name: `Watchlist #${watchlists.length + 1}`,
    //     enabled: true,
    //     items: [],
    //   },
    // ]);

    console.log("adding watchlist")
  };

  const toggleWatchlist = (id) => {
    setWatchlists((prev) =>
      prev.map((w) =>
        w.id === id ? { ...w, enabled: !w.enabled } : w
      )
    );
  };

  return (
    <>
      <Header />
      {showAddWatchlistPopup && (
        <AddWatchlistPopup onClose={() => setShowAddPopup(false)} />
      )}
      <WatchlistGrid
        watchlists={watchlists}
        onAdd={addWatchlist}
        onToggle={toggleWatchlist}
      />
    </>
  );
}
