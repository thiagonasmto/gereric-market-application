import React from 'react';
import ProductList from '../services/ProductList';
import styles from './styles/Products.module.css';

const Products: React.FC = () => {
  return (
    <div className={styles.products}>
      <ProductList />
    </div>
  );
};

export default Products;
