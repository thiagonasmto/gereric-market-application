import React, { useEffect, useState } from 'react';
import Order from '../components/Order';
import ProductCard from '../components/ProductCard';
import { API_BASE_URL } from '../services/api';

type Product = {
  id: string;
  name: string;
  description: string;
  price: number;
};

type SelectedProduct = {
  productid: string;
  name: string;
  price: number;
  quantity: number;
};

const ProductSelectionAndOrder: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]);
  const [selectedProducts, setSelectedProducts] = useState<SelectedProduct[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(`${API_BASE_URL}/products/`)
      .then((res) => res.json())
      .then((data) => {
        setProducts(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error('Erro ao buscar produtos:', err);
        setLoading(false);
      });
  }, []);

  const handleProductClick = (product: Product) => {
    setSelectedProducts((prev) => {
      const exists = prev.find((item) => item.productid === product.id);
      if (exists) return prev;
      return [...prev, {
        productid: product.id,
        name: product.name,
        price: product.price,
        quantity: 1,
      }];
    });
  };

  const handleQuantityChange = (productid: string, quantity: number) => {
    setSelectedProducts((prev) => {
      if (quantity <= 0) {
        return prev.filter((item) => item.productid !== productid);
      }

      return prev.map((item) =>
        item.productid === productid ? { ...item, quantity } : item
      );
    });
  };

  if (loading) return <p>Carregando produtos...</p>;

  return (
    <div style={styles.pageContainer}>
      <div style={styles.productsGrid}>
        {products.map((product) => (
          <div key={product.id} onClick={() => handleProductClick(product)}>
            <ProductCard
              name={product.name}
              description={product.description}
              price={product.price}
            />
          </div>
        ))}
      </div>

      <div style={styles.orderPanel}>
        <Order
          selectedProducts={selectedProducts}
          onQuantityChange={handleQuantityChange}
        />
      </div>
    </div>
  );
};

const styles: { [key: string]: React.CSSProperties } = {
  pageContainer: {
    display: 'flex',
    justifyContent: 'space-between',
    padding: '30px 50px',
    gap: 40,
    backgroundColor: '#f9f9f9',
    width: '100%',
  },
  productsGrid: {
    display: 'grid',
    gridTemplateColumns: 'repeat(auto-fill, minmax(250px, 1fr))',
    gap: 24,
    flex: 1,
  },
  orderPanel: {
    width: 320,
    position: 'sticky',
    top: 30,
    alignSelf: 'flex-start',
  },
};

export default ProductSelectionAndOrder;
