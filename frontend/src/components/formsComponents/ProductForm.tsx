import React, { useState } from 'react';
import styles from '../styles/ProductForm.module.css';
import { API_BASE_URL } from '../../services/api';

const ProductForm: React.FC = () => {
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    quantity: 0,
    price: 0,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;

    if (name === 'description') {
      if (value.length > 100) return;
      setFormData(prev => ({ ...prev, [name]: value }));
    } else if (name === 'quantity' || name === 'price') {
      const numberValue = Number(value);
      if (!isNaN(numberValue)) {
        setFormData(prev => ({ ...prev, [name]: numberValue }));
      }
    } else {
      setFormData(prev => ({ ...prev, [name]: value }));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = localStorage.getItem("authToken");
    try {
      const response = await fetch(`${API_BASE_URL}/products/`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) throw new Error('Erro na requisição');

      alert('Produto cadastrado com sucesso!');
      setFormData({ name: '', description: '', quantity: 0, price: 0 });
    } catch (error) {
      alert('Erro ao cadastrar produto.');
    }
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.title}>Cadastro de Produto</h2>
      <form onSubmit={handleSubmit} className={styles.form}>
        <div className={styles.row}>
          <div className={styles.cell}>
            <label htmlFor="name" className={styles.label}>Nome:</label>
            <input
              id="name"
              name="name"
              type="text"
              value={formData.name}
              onChange={handleChange}
              required
              className={styles.input}
            />
          </div>

          <div className={styles.cell}>
            <label htmlFor="quantity" className={styles.label}>Quantidade:</label>
            <input
              id="quantity"
              name="quantity"
              type="number"
              min={0}
              value={formData.quantity}
              onChange={handleChange}
              required
              className={styles.input}
            />
          </div>

          <div className={styles.cell}>
            <label htmlFor="price" className={styles.label}>Preço:</label>
            <input
              id="price"
              name="price"
              type="number"
              min={0}
              value={formData.price}
              onChange={handleChange}
              required
              className={styles.input}
            />
          </div>
        </div>

        <div className={styles.fullWidth}>
          <label htmlFor="description" className={styles.label}>
            Descrição (máx. 100 caracteres):
          </label>
          <textarea
            id="description"
            name="description"
            value={formData.description}
            onChange={handleChange}
            required
            className={`${styles.input} ${styles.textarea}`}
          />
          <small className={styles.charCount}>
            {formData.description.length} / 100 caracteres
          </small>
        </div>

        <div className={styles.buttonWrapper}>
          <button type="submit" className={styles.button}>Cadastrar</button>
        </div>
      </form>
    </div>
  );
};

export default ProductForm;
