import React, { useEffect, useState } from 'react';
import { Table, Typography, Card, Tag, Popover, Button, Divider } from 'antd';
import styles from '../styles/OrdersTableInProgress.module.css';
import { API_BASE_URL } from '../../services/api';

const { Title } = Typography;

type Product = {
  id: string;
  name: string;
  quantity: number;
};

type OrderInProgress = {
  id: string;
  client: {
    email: string;
    id: string;
  };
  products: Product[];
  status: 'Em andamento';
  totalPrice: number;
};

export const OrdersTableInProgress: React.FC = () => {
  const token = localStorage.getItem('authToken');
  const [orders, setOrders] = useState<OrderInProgress[]>([]);

  useEffect(() => {
    fetch(`${API_BASE_URL}/services/ordes-in-progress`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      }
    })
      .then(res => res.json())
      .then(data => setOrders(data))
      .catch(err => console.error('Erro ao buscar pedidos em andamento:', err));
  }, []);

  const columns = [
    {
      title: 'Email do Cliente',
      dataIndex: ['client', 'email'],
      key: 'clientEmail',
    },
    {
      title: 'Produtos',
      dataIndex: 'products',
      key: 'products',
      render: (products: Product[] | null) => {
        if (!products || products.length === 0) return 'Nenhum produto';

        const content = (
          <div style={{ padding: '0 10px' }}>
            {products.map((p, index) => (
              <div key={index} style={{ padding: '5px 0' }}>
                <div><strong>Nome:</strong> {p.name ?? 'Desconhecido'}</div>
                <div><strong>Qtd:</strong> {p.quantity}</div>
                {index < products.length - 1 && <Divider style={{ margin: '8px 0' }} />}
              </div>
            ))}
          </div>
        );

        return (
          <Popover content={content} title="Produtos do Pedido" trigger="click">
            <Button type="link">Ver Produtos ({products.length})</Button>
          </Popover>
        );
      },
    },
    {
      title: 'Preço Total',
      dataIndex: 'totalPrice',
      key: 'totalPrice',
      render: (price: number) => `R$ ${price.toFixed(2)}`,
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      render: (status: string) => (
        <Tag className={styles.tagYellow}>
          {status}
        </Tag>
      ),
    },
  ];

  return (
    <Card className={styles.card}>
      <Title level={3} className={styles.title}>Pedidos em Andamento</Title>
      <Table
        dataSource={orders}
        columns={columns}
        rowKey={order => order.id}
        bordered
      />
    </Card>
  );
};
