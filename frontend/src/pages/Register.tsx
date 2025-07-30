import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import styles from './styles/Register.module.css';
import { API_BASE_URL } from '../services/api';

const Register: React.FC = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!name || !email || !password) {
      alert('Por favor, preencha todos os campos.');
      return;
    }

    setLoading(true);

    try {
      const response = await fetch(`${API_BASE_URL}/clients/`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        alert(errorData.message || 'Erro ao cadastrar cliente.');
        return;
      }

      alert('Cadastro realizado com sucesso!');
      navigate('/login');
    } catch (error) {
      console.error('Erro ao cadastrar:', error);
      alert('Erro inesperado. Tente novamente.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={styles.container}>
      <form className={styles.form} onSubmit={handleRegister}>
        <h2 className={styles.title}>Cadastro</h2>
        <hr className={styles.divider} />
        <p className={styles.subtitle}>Crie sua conta para continuar</p>

        <label htmlFor="name" className={styles.label}>Nome</label>
        <input
          id="name"
          type="text"
          placeholder="Seu nome"
          value={name}
          onChange={e => setName(e.target.value)}
          className={styles.input}
        />

        <label htmlFor="email" className={styles.label}>Email</label>
        <input
          id="email"
          type="email"
          placeholder="Seu email"
          value={email}
          onChange={e => setEmail(e.target.value)}
          className={styles.input}
        />

        <label htmlFor="password" className={styles.label}>Senha</label>
        <input
          id="password"
          type="password"
          placeholder="Sua senha"
          value={password}
          onChange={e => setPassword(e.target.value)}
          className={styles.input}
        />

        <button type="submit" disabled={loading} className={styles.button}>
          {loading ? 'Cadastrando...' : 'CADASTRAR'}
        </button>

        <p className={styles.loginLink}>
          Já tem conta?{' '}
          <a href="/login" className={styles.loginAnchor}>
            Faça login
          </a>
        </p>
      </form>
    </div>
  );
};

export default Register;
