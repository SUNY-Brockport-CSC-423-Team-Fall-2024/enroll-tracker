"use client";

import { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import styles from "../styles.module.css";

interface User {
  username: string;
  id: number;
  first_name: string;
  last_name: string;
  phone_number: string;
  email: string;
  office?: string; // Only for teachers
}

export default function UserProfile() {
  const [userData, setUserData] = useState<User | null>(null);
  const [formData, setFormData] = useState<Partial<User>>({});
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();
  const params = useParams();
  const username = params.username as string;

  useEffect(() => {
    const fetchUserData = async () => {
      if (!username) return;
      try {
        const userType = await getUserType(username);
        const url = `http://localhost:8002/api/${userType}/${username}`;
        const response = await fetch(url);
        if (!response.ok) {
          throw new Error("Error fetching user data");
        }
        const data: User = await response.json();
        setUserData(data);
        setFormData(data); // Initialize form data with user data
      } catch (err) {
        setError(err instanceof Error ? err.message : "An unknown error occurred");
      }
    };
    fetchUserData();
  }, [username]);

  const getUserType = async (username: string) => {
    const studentResponse = await fetch(`http://localhost:8002/api/students/${username}`);
    if (studentResponse.ok) return "students";
    const teacherResponse = await fetch(`http://localhost:8002/api/teachers/${username}`);
    if (teacherResponse.ok) return "teachers";
    throw new Error("User not found");
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSave = async () => {
    if (!username || !userData) return;
    const userType = userData.office ? "teachers" : "students";
    const url = `http://localhost:8002/api/${userType}/${username}`;

    try {
      const response = await fetch(url, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });
      if (!response.ok) {
        throw new Error("Error updating user data");
      }
      alert("User data updated successfully!");
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred while updating.");
    }
  };

  return (
    <div className={styles.users_root}>
      <header className={styles.header}>
        <h1>Modify User</h1>
        <button onClick={() => router.push("/users")} className={styles.right_button}>
          Back
        </button>
      </header>

      {error && <p className={styles.error}>Error: {error}</p>}

      <div className={styles.form_container}>
        <div className={styles.field_container}>
          <label>First Name</label>
          <input
            type="text"
            name="first_name"
            placeholder={userData?.first_name}
            value={formData.first_name || ""}
            onChange={handleChange}
            className={styles.input_field}
          />
        </div>
        <div className={styles.field_container}>
          <label>Last Name</label>
          <input
            type="text"
            name="last_name"
            placeholder={userData?.last_name}
            value={formData.last_name || ""}
            onChange={handleChange}
            className={styles.input_field}
          />
        </div>
        <div className={styles.field_container}>
          <label>Phone Number</label>
          <input
            type="text"
            name="phone_number"
            placeholder={userData?.phone_number}
            value={formData.phone_number || ""}
            onChange={handleChange}
            className={styles.input_field}
          />
        </div>
        <div className={styles.field_container}>
          <label>Email</label>
          <input
            type="text"
            name="email"
            placeholder={userData?.email}
            value={formData.email || ""}
            onChange={handleChange}
            className={styles.input_field}
          />
        </div>
        {userData?.office && (
          <div className={styles.field_container}>
            <label>Office</label>
            <input
              type="text"
              name="office"
              placeholder={userData.office}
              value={formData.office || ""}
              onChange={handleChange}
              className={styles.input_field}
            />
          </div>
        )}
      </div>

      <button onClick={handleSave} className={styles.centered_button}>
        Save
      </button>
    </div>
  );
}
