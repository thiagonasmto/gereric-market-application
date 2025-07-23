import React, { useState } from 'react';
import styles from './styles/Footer.module.css';

const Footer: React.FC = () => {
  const [inputText, setInputText] = useState('');
  const [vogalResult, setVogalResult] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8081/services/find-vogal', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ input: inputText }),
      });

      if (response.ok) {
        const data = await response.json();
        setVogalResult(data.vogal);
      } else {
        setVogalResult('Erro na requisição');
      }
    } catch (error) {
      setVogalResult('Erro ao conectar com o servidor');
    }
  };

  return (
    <footer className={styles.footer}>
      <form onSubmit={handleSubmit} className={styles.vogalForm}>
        <input
          type="text"
          value={inputText}
          onChange={(e) => setInputText(e.target.value)}
          placeholder="Digite a string"
          className={styles.vogalInput}
        />
        <button type="submit" className={styles.vogalButton}>
          Enviar
        </button>
      </form>
      {vogalResult && (
        <div className={styles.vogalResult}>
          Vogal retornada: <strong>{vogalResult}</strong>
        </div>
      )}
      <div>Desenvolvido por Thiago Lopes</div>
    </footer>
  );
};

export default Footer;
