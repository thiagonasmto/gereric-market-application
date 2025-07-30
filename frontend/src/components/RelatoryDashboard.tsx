import React, { useState } from 'react';
import { SummaryTable } from './tablesComponents/SummaryTable';
import { OrdersTableInProgress } from './tablesComponents/OrdersTableInProgress';
import { RankClientsTable } from './tablesComponents/RankClientsTable';
import styles from './styles/RelatoryDasboard.module.css'; 
import { API_BASE_URL } from '../services/api';

const RelatoryDashboard: React.FC = () => {
  const [activeTable, setActiveTable] = useState<'none' | 'inProgress' | 'rank'>('none');

  const renderTable = () => {
    switch (activeTable) {
      case 'inProgress':
        return <OrdersTableInProgress />;
      case 'rank':
        return <RankClientsTable />;
      default:
        return null;
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.summaryWrapper}>
        <SummaryTable />
      </div>
      <div className={styles.buttonsWrapper}>
        <button
          onClick={() => setActiveTable('inProgress')}
          className={styles.buttonYellow}
        >
          Pedidos em Andamento
        </button>
        <button
          onClick={() => setActiveTable('rank')}
          className={styles.buttonBlue}
        >
          Ranking de Clientes
        </button>
        <button
          onClick={async () => {
            const authToken = localStorage.getItem('authToken');
            if (!authToken) {
              alert('Token não encontrado. Faça login novamente.');
              return;
            }

            try {
              const response = await fetch(`${API_BASE_URL}/services/generate-excel`, {
                method: 'GET',
                headers: {
                  'Authorization': `Bearer ${authToken}`
                }
              });

              if (!response.ok) {
                throw new Error('Erro ao baixar o relatório');
              }

              const blob = await response.blob();
              const url = window.URL.createObjectURL(blob);
              const link = document.createElement('a');
              link.href = url;
              link.setAttribute('download', 'relatorio-gestao-de-vendas.xlsx');
              document.body.appendChild(link);
              link.click();
              link.remove();
              window.URL.revokeObjectURL(url);
            } catch (err) {
              console.error('Erro no download:', err);
              alert('Erro ao fazer o download do relatório.');
            }
          }}
          className={styles.buttonGreen}
        >
          Download Relatórios
        </button>
      </div>
      {renderTable()}
    </div>
  );
};

export default RelatoryDashboard;
