import React, { useEffect, useState } from 'react';
import { Table, Typography, Card } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import styles from '../styles/RanckClientsTable.module.css';
import { API_BASE_URL } from '../../services/api';

const { Title } = Typography;

type RankedClient = {
  createdAt: string;
  email: string;
  ordersQuantity: number;
};

export const RankClientsTable: React.FC = () => {
  const token = localStorage.getItem('authToken');
  const [clients, setClients] = useState<RankedClient[]>([]);

  useEffect(() => {
    fetch(`${API_BASE_URL}/services/rank-clients`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      }
    })
      .then(res => res.json())
      .then(data => setClients(data))
      .catch(err => console.error('Erro ao buscar ranking de clientes:', err));
  }, []);

  const columns: ColumnsType<RankedClient> = [
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
    },
    {
      title: 'Data de Cadastro',
      dataIndex: 'createdAt',
      key: 'createdAt',
      render: (date: string) =>
        new Date(date).toLocaleDateString('pt-BR', {
          day: '2-digit',
          month: '2-digit',
          year: 'numeric',
        }),
    },
    {
      title: 'Total de Pedidos',
      dataIndex: 'ordersQuantity',
      key: 'ordersQuantity',
      sorter: (a, b) => a.ordersQuantity - b.ordersQuantity,
      defaultSortOrder: 'descend',
    },
  ];

  return (
    <Card className={styles.card}>
      <Title level={3} className={styles.title}>Ranking de Clientes</Title>
      <Table<RankedClient>
        dataSource={clients}
        columns={columns}
        rowKey={(record) => record.email}
        bordered
      />
    </Card>
  );
};
