import React, { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { jwtDecode } from "jwt-decode";
import styles from './styles/Login.module.css';

interface DecodedToken {
  clientid: string;
  role: string;
  exp: number;
}

const Login: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isAdmin, setIsAdmin] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8081/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password, isAdmin }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        alert(errorData.message || 'Erro ao fazer login.');
        return;
      }

      const data = await response.json();
      const token = data.token;
      localStorage.setItem('authToken', token);

      const decoded = jwtDecode<DecodedToken>(token);

      if (decoded.role === 'admin') {
        navigate('/relatory');
      } else {
        const from = (location.state as { from?: string })?.from || '/products';
        navigate(from);
      }

    } catch (error) {
      console.error('Erro ao fazer login:', error);
      alert('Erro inesperado. Tente novamente.');
    }
  };

  return (
    <div className={styles.container}>
      <form className={styles.form} onSubmit={handleLogin}>
        <h2 className={styles.title}>Login</h2>
        <hr className={styles.divider} />
        <p className={styles.subtitle}>Welcome to the Generic Market!</p>

        <label htmlFor="email" className={styles.label}>Email</label>
        <input
          id="email"
          type="text"
          placeholder="Enter your email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className={styles.input}
        />

        <label htmlFor="password" className={styles.label}>Password</label>
        <input
          id="password"
          type="password"
          placeholder="Enter your password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className={styles.input}
        />

        <label className={styles.checkboxLabel}>
          <input
            type="checkbox"
            checked={isAdmin}
            onChange={() => setIsAdmin(!isAdmin)}
          />
          Login como Admin
        </label>

        <button type="submit" className={styles.button}>
          LOGIN
        </button>

        <p className={styles.registerLink}>
          New to Market?{' '}
          <a href="/register" className={styles.anchor}>
            Register Here
          </a>
        </p>
      </form>
    </div>
  );
};

export default Login;
