import React, { useEffect, useState } from 'react';
import { Table, Typography, Card } from 'antd';
import styles from '../styles/SummaryTable.module.css';

const { Title } = Typography;

type Summary = {
  totalSalesMade: number;
  invoicedSmount: number;
  totalOrdersSold: number;
};

export const SummaryTable: React.FC = () => {
  const token = localStorage.getItem('authToken');
  const [summary, setSummary] = useState<Summary | null>(null);

  useEffect(() => {
    fetch('http://localhost:8081/services/summary', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      }
    })
    .then(res => res.json())
    .then(data =>
      setSummary({
        totalSalesMade: data.totalSalesMade,
        invoicedSmount: data.invoicedSmount,
        totalOrdersSold: data.totalOrdersSold,
      })
    )
    .catch(err => console.error('Erro ao buscar resumo:', err));
}, []);

const columns = [
  {
    title: 'Total de Vendas Realizadas',
    dataIndex: 'totalSalesMade',
    key: 'totalSalesMade',
  },
  {
    title: 'Valor Faturado (R$)',
    dataIndex: 'invoicedSmount',
    key: 'invoicedSmount',
    render: (value: number) => `R$ ${value.toFixed(2)}`,
  },
  {
    title: 'Total de Pedidos Vendidos',
    dataIndex: 'totalOrdersSold',
    key: 'totalOrdersSold',
  },
];

return (
  <Card className={styles.card}>
    <Title level={3} className={styles.title}>Resumo Geral</Title>
    <Table
      dataSource={summary ? [summary] : []}
      columns={columns}
      pagination={false}
      rowKey={() => 'summary-row'}
      bordered
    />
  </Card>
);
};
