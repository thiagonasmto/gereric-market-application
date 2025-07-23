import { Routes, Route } from 'react-router-dom';
import Home from '../pages/Home';
import Login from '../pages/Login';
import Products from '../pages/Products';
import Register from '../pages/Register';
import Relatory from '../pages/Relatory';

import Header from '../components/Header';
import Footer from '../components/Footer';

const DefaultLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => (
  <div style={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
    <Header />
    {children}
    <Footer />
  </div>
);

export default function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<DefaultLayout><Home /></DefaultLayout>} />
      <Route path="/login" element={<DefaultLayout><Login /></DefaultLayout>} />
      <Route path="/register" element={<DefaultLayout><Register /></DefaultLayout>} />
      <Route path="/products" element={<DefaultLayout><Products /></DefaultLayout>} />

      <Route path="/relatory" element={<Relatory />} />
    </Routes>
  );
}
