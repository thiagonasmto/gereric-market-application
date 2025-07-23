import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Importa o hook
import { FaUserCircle, FaChartPie, FaSignOutAlt } from 'react-icons/fa';
import RelatoryDashboard from '../components/RelatoryDashboard';
import AdminForm from '../components/formsComponents/AdminForm';
import ProductForm from '../components/formsComponents/ProductForm';
import { OrdersTable } from '../components/tablesComponents/OrdersTable';
import styles from './styles/Relatory.module.css';

const Relatory: React.FC = () => {
  const [activeComponent, setActiveComponent] = useState<'dashboard' | 'admins' | 'products' | 'orders'>('dashboard');
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    navigate('/login');
  };

  const renderContent = () => {
    switch (activeComponent) {
      case 'admins':
        return <AdminForm />;
      case 'products':
        return <ProductForm />;
      case 'orders':
        return <OrdersTable />;
      case 'dashboard':
        return <RelatoryDashboard />;
      default:
        return <RelatoryDashboard />;
    }
  };

  return (
    <div className={styles.container}>
      <aside className={styles.sidebar}>
        <div>
          <div className={styles.headerUser}>
            <FaUserCircle size={20} className={styles.iconSpacing} />
            <span>Nome do Adm</span>
          </div>

          <ul className={styles.navList}>
            <li onClick={() => setActiveComponent('admins')}>• Gerenciar Administradores</li>
            <li onClick={() => setActiveComponent('products')}>• Gerenciar Produtos</li>
            <li onClick={() => setActiveComponent('orders')}>• Gerenciar Pedidos</li>

            <div className={styles.relatoryTitle}>Relatórios</div>
            <li className={styles.dashboardItem} onClick={() => setActiveComponent('dashboard')}>
              <FaChartPie size={12} className={styles.iconSpacing} />
              Overview
            </li>
          </ul>
        </div>

        <div className={styles.logout} onClick={handleLogout}>
          <FaSignOutAlt size={14} className={styles.iconSpacing} />
          Sair
        </div>
      </aside>

      <main className={styles.mainContent}>
        {renderContent()}
      </main>
    </div>
  );
};

export default Relatory;
