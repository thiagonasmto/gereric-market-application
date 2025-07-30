import React, { useState } from 'react';
import styles from '../styles/AdminForm.module.css';
import { API_BASE_URL } from '../../services/api';

const AdminForm: React.FC = () => {
  const [formData, setFormData] = useState({ name: '', email: '', password: '' });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = localStorage.getItem("authToken");
    try {
      const response = await fetch(`${API_BASE_URL}/adms/`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(formData)
      });

      if (!response.ok) throw new Error('Erro na requisição');

      alert('Administrador cadastrado com sucesso!');
      setFormData({ name: '', email: '', password: '' });
    } catch (error) {
      alert('Erro ao cadastrar administrador.');
    }
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.title}>Cadastro de Administrador</h2>
      <form onSubmit={handleSubmit} className={styles.form}>
        <input
          name="name"
          placeholder="Nome"
          value={formData.name}
          onChange={handleChange}
          required
          className={styles.input}
        />
        <input
          name="email"
          type="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange}
          required
          className={styles.input}
        />
        <input
          name="password"
          type="password"
          placeholder="Senha"
          value={formData.password}
          onChange={handleChange}
          required
          className={styles.input}
        />
        <button
          type="submit"
          className={styles.button}
        >
          Cadastrar
        </button>
      </form>
    </div>
  );
};

export default AdminForm;
