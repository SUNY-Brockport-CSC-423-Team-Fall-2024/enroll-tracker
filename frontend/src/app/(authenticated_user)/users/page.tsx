"use client";

import { useState, useEffect } from "react";
import styles from "./styles.module.css";
import { useAuth } from "@/app/providers/auth-provider";
import { useRouter } from "next/navigation";

interface User {
  username: string;
  id: number;
  first_name: string;
  last_name: string;
}

export default function Users() {
  const router = useRouter();
  const { username } = useAuth();
  const [selectedButton, setSelectedButton] = useState<string>("Students");
  const [students, setStudents] = useState<User[]>([]);
  const [teachers, setTeachers] = useState<User[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [userType, setUserType] = useState<"student" | "teacher">("student");
  const [formData, setFormData] = useState<any>({}); // Form data for adding a user

  const handleButtonClick = (button: string) => {
    setSelectedButton(button);
  };

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        let url = "";
        let data: User[] = [];

        if (selectedButton === "Students") {
          // Fetch students data
          url = `http://localhost:8002/api/students`;
          const response = await fetch(url);
          if (!response.ok) {
            throw new Error("Error fetching students");
          }
          data = await response.json();
          setStudents(data);
        } else if (selectedButton === "Teachers") {
          // Fetch teachers data
          url = `http://localhost:8002/api/teachers`;
          const response = await fetch(url);
          if (!response.ok) {
            throw new Error("Error fetching teachers");
          }
          data = await response.json();
          setTeachers(data);
        }
      } catch (err) {
        setError(err instanceof Error ? err.message : "An unknown error occurred");
      }
    };

    fetchUsers();
  }, [selectedButton]);

  const goToUserProfile = (username: string) => {
    const profileUsername = username || "student123";
    router.push(`/users/${profileUsername}`);
  };

  const handleUserTypeChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setUserType(e.target.value as "student" | "teacher");
  };

  const handleFormChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const url =
        userType === "student"
          ? "http://localhost:8002/api/students"
          : "http://localhost:8002/api/teachers";

      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error(`Error creating ${userType}`);
      }

      const result = await response.json();
      alert(`${userType.charAt(0).toUpperCase() + userType.slice(1)} created successfully!`);

      setFormData({});
    } catch (err) {
      alert(err instanceof Error ? err.message : "An unknown error occurred");
    }
  };

  const buttons = ["Students", "Teachers", "Add User"];

  return (
    <div className={styles.users_root}>
      <header className={styles.header}>
        <h1>Users</h1>
        <button className={styles.right_button}>{username}</button>
      </header>

      <nav className={styles.nav_bar}>
        {buttons.map((button) => (
          <button
            key={button}
            onClick={() => handleButtonClick(button)}
            className={`${styles.nav_button} ${selectedButton === button ? styles.selected : ""}`}
          >
            {button}
          </button>
        ))}
      </nav>

      <div className={styles.scroll_list}>
        {selectedButton === "Students" && (
          <>
            {error && <p className={styles.error}>Error: {error}</p>}

            <div className={styles.header_bar}>
              <span className={styles.column_header}>Student ID</span>
              <span className={styles.column_header}>First Name</span>
              <span className={styles.column_header}>Last Name</span>
            </div>

            {students.length > 0 ? (
              students.map((student, index) => (
                <div
                  key={index}
                  className={styles.list_item}
                  onClick={() => goToUserProfile(student.username)}
                  style={{ cursor: "pointer" }}
                >
                  <span className={styles.username}>{student.id}</span>
                  <span className={styles.first_name}>{student.first_name}</span>
                  <span className={styles.last_name}>{student.last_name}</span>
                </div>
              ))
            ) : (
              <p>No students found.</p>
            )}
          </>
        )}
        {selectedButton === "Teachers" && (
          <>
            {error && <p className={styles.error}>Error: {error}</p>}

            <div className={styles.header_bar}>
              <span className={styles.column_header}>Teacher ID</span>
              <span className={styles.column_header}>First Name</span>
              <span className={styles.column_header}>Last Name</span>
            </div>

            {teachers.length > 0 ? (
              teachers.map((teacher, index) => (
                <div
                  key={index}
                  className={styles.list_item}
                  onClick={() => goToUserProfile(teacher.username)}
                  style={{ cursor: "pointer" }}
                >
                  <span className={styles.username}>{teacher.id}</span>
                  <span className={styles.first_name}>{teacher.first_name}</span>
                  <span className={styles.last_name}>{teacher.last_name}</span>
                </div>
              ))
            ) : (
              <p>No teachers found.</p>
            )}
          </>
        )}
        {selectedButton === "Add User" && (
          <div className={styles.form_container}>
            <form onSubmit={handleSubmit}>
              <div className={styles.field_container}>
                <label>User Type</label>
                <select
                  value={userType}
                  onChange={handleUserTypeChange}
                  className={styles.input_field}
                >
                  <option value="student">Student</option>
                  <option value="teacher">Teacher</option>
                </select>
              </div>

              <div className={styles.field_container}>
                <label>Username</label>
                <input
                  type="text"
                  name="username"
                  value={formData.username || ""}
                  onChange={handleFormChange}
                  className={styles.input_field}
                  required
                />
              </div>

              <div className={styles.field_container}>
                <label>Password</label>
                <input
                  type="password"
                  name="password"
                  value={formData.password || ""}
                  onChange={handleFormChange}
                  className={styles.input_field}
                  required
                />
              </div>

              <div className={styles.field_container}>
                <label>First Name</label>
                <input
                  type="text"
                  name="first_name"
                  value={formData.first_name || ""}
                  onChange={handleFormChange}
                  className={styles.input_field}
                  required
                />
              </div>

              <div className={styles.field_container}>
                <label>Last Name</label>
                <input
                  type="text"
                  name="last_name"
                  value={formData.last_name || ""}
                  onChange={handleFormChange}
                  className={styles.input_field}
                  required
                />
              </div>

              <div className={styles.field_container}>
                <label>Phone Number</label>
                <input
                  type="text"
                  name="phone_number"
                  value={formData.phone_number || ""}
                  onChange={handleFormChange}
                  className={styles.input_field}
                  required
                />
              </div>

              <div className={styles.field_container}>
                <label>Email</label>
                <input
                  type="text"
                  name="email"
                  value={formData.email || ""}
                  onChange={handleFormChange}
                  className={styles.input_field}
                  required
                />
              </div>

              {userType === "teacher" && (
                <div className={styles.field_container}>
                  <label>Office</label>
                  <input
                    type="text"
                    name="office"
                    value={formData.office || ""}
                    onChange={handleFormChange}
                    className={styles.input_field}
                    required
                  />
                </div>
              )}

              <button type="submit" className={styles.centered_button}>
                Save User
              </button>
            </form>
          </div>
        )}
      </div>
    </div>
  );
}
