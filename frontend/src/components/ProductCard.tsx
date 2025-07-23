import React from 'react';
import styles from './styles/ProductCard.module.css';

type ProductCardProps = {
  name: string;
  description: string;
  price: number;
};

const ProductCard: React.FC<ProductCardProps> = ({ name, description, price }) => {
  return (
    <div className={styles.card}>
      <div className={styles.imagePlaceholder} />
      <div className={styles.content}>
        <h3 className={styles.title}>{name}</h3>
        <p className={styles.description}>{description}</p>
        <p className={styles.price}>R$ {price.toLocaleString()}</p>
      </div>
    </div>
  );
};

export default ProductCard;
