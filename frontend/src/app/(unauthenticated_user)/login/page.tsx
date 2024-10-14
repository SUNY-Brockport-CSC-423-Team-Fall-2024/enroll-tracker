'use client'

import styles from '../styles.module.css'

export default function Login() {

    const handleLogin = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log('hi');
    }

    return (
        <div className={styles.main_content}>
            <div className={styles.login_box}>
                <h2>Login</h2>
                <form onSubmit={handleLogin} className={styles.login_form}>
                    <div className={styles.login_form_field}>
                        <label htmlFor="username">Username</label>
                        <input id="username_input" type="text" name="username" />
                    </div>
                    <div className={styles.login_form_field}>
                        <label htmlFor="password">Password</label>
                        <input id="password_input" type="password" name="password" />
                    </div>
                    <div className={styles.login_form_field}>
                        <input id={styles.login_submit} type="submit" value="Login" /> 
                    </div>
                </form>
            </div>
        </div>
    )
}
