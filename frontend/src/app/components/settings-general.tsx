"use client";

import React, { useState } from "react";
import "./styles.module.css";

const SettingsGeneral = () => {
  return (
    <div className="setting-gen-box">
      <p>This is placeholder and not integrated</p>
      <div className="settings-user-email">
        <div>
          <label>Username</label>
          <input type="text" placeholder="uname00" readOnly />
        </div>
        <br />
        <div>
          <label>Email</label>
          <input type="text" placeholder="user00@brockport.edu" readOnly />
        </div>
      </div>
    </div>
  );
};

export default SettingsGeneral;
