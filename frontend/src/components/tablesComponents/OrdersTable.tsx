import React, { useEffect, useState } from 'react';
import { Table, Typography, Card, Select, Popover, Button, Divider } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import styles from '../styles/OrdersTable.module.css';
import { API_BASE_URL } from '../../services/api';

const { Title } = Typography;
const { Option } = Select;

type Order = {
  id: string;
  client: {
    email: string;
    id: string;
  };
  products: { id: string; name: string; quantity: number }[] | null;
  status: 'Em andamento' | 'Finalizado' | 'Cancelado';
  totalPrice: number;
};

type Product = {
  id: string;
  name: string;
  quantity: number;
};

const statusOptions: Order['status'][] = ['Em andamento', 'Finalizado', 'Cancelado'];

export const OrdersTable: React.FC = () => {
  const token = localStorage.getItem('authToken');
  const [orders, setOrders] = useState<Order[]>([]);

  useEffect(() => {
    fetch(`${API_BASE_URL}/orders/`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      }
    })
      .then(res => res.json())
      .then(data => setOrders(data))
      .catch(err => console.error('Erro ao buscar pedidos:', err));
  }, []);

  const handleStatusChange = async (id: string, newStatus: Order['status']) => {
    try {
      const response = await fetch(`${API_BASE_URL}/orders/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({ status: newStatus }),
      });

      if (!response.ok) throw new Error('Erro ao atualizar status');

      setOrders(prev =>
        prev.map(order =>
          order.id === id ? { ...order, status: newStatus } : order
        )
      );
    } catch (error) {
      console.error(error);
      alert('Erro ao atualizar status do pedido.');
    }
  };

  const columns: ColumnsType<Order> = [
    {
      title: 'Email do Cliente',
      dataIndex: ['client', 'email'],
      key: 'email',
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
      render: (value: number) => `R$ ${value.toFixed(2)}`,
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      render: (status, record) => (
        <Select
          value={status}
          className={styles.select}
          onChange={(value) => handleStatusChange(record.id, value)}
        >
          {statusOptions.map(option => (
            <Option key={option} value={option}>
              {option}
            </Option>
          ))}
        </Select>
      ),
    },
  ];

  return (
    <Card className={styles.card}>
      <Title level={3} className={styles.title}>Gerenciar Pedidos</Title>
      <Table<Order>
        dataSource={orders}
        columns={columns}
        rowKey={order => order.id}
        bordered
      />
    </Card>
  );
};
