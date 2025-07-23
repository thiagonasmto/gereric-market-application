import React from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import styles from './styles/Header.module.css';

const Header: React.FC = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const isAuthenticated = Boolean(localStorage.getItem('authToken'));

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    navigate('/login');
  };

  if (location.pathname === '/login') {
    return (
      <header className={styles.header}>
        <nav>
          <a href="/products" className={styles.navLink}>Produtos</a>
          <a href="/" className={styles.navLink}>Home</a>
          {!isAuthenticated ? "" : (
            <button onClick={handleLogout} className={styles.navLink} style={{ background: 'none', border: 'none', cursor: 'pointer' }}>
              Logout
            </button>
          )}
        </nav>
      </header>
    );
  }

  if (location.pathname === '/products') {
    return (
      <header className={styles.header}>
        <nav>
          <a href="/" className={styles.navLink}>Home</a>
          {!isAuthenticated ? (
            <a href="/login" className={styles.navLink}>Login</a>
          ) : (
            <button onClick={handleLogout} className={styles.navLink} style={{ background: 'none', border: 'none', cursor: 'pointer' }}>
              Logout
            </button>
          )}
        </nav>
      </header>
    );
  }

  if (location.pathname === '/register') {
    return (
      <header className={styles.header}>
      <nav>
        <a href="/products" className={styles.navLink}>Produtos</a>
        <a href="/" className={styles.navLink}>Home</a>
        {!isAuthenticated ? (
            <a href="/login" className={styles.navLink}>Login</a>
          ) : (
            <button onClick={handleLogout} className={styles.navLink} style={{ background: 'none', border: 'none', cursor: 'pointer' }}>
              Logout
            </button>
          )}
      </nav>
    </header>
    );
  }

  return (
    <header className={styles.header}>
      <nav>
        <a href="/products" className={styles.navLink}>Produtos</a>
        {!isAuthenticated ? (
            <a href="/login" className={styles.navLink}>Login</a>
          ) : (
            <button onClick={handleLogout} className={styles.navLink} style={{ background: 'none', border: 'none', cursor: 'pointer' }}>
              Logout
            </button>
          )}
      </nav>
    </header>
  );
};

export default Header;
