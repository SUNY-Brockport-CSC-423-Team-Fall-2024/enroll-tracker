"use client";

import React, { useState } from "react";
import "./styles.module.css";
import SettingsGeneral from "./settings-general";
import SettingsAuth from "./settings-auth";

const SettingsTabBar = () => {
  const [view, setView] = useState("gen");

  const handleGenClick = () => {
    setView("gen");
  };

  const handleAuthClick = () => {
    setView("auth");
  };

  return (
    <div>
      <div className="styles.settings-nav-bar">
        <button onClick={handleGenClick} className={view === "gen" ? "selected" : "not-selected"}>
          General
        </button>
        <button onClick={handleAuthClick} className={view === "auth" ? "selected" : "not-selected"}>
          Authentication
        </button>
      </div>

      {view === "gen" && <SettingsGeneral />}

      {view === "auth" && <SettingsAuth />}
    </div>
  );
};

export default SettingsTabBar;
