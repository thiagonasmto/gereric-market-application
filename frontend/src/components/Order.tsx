import React, { useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';
import styles from './styles/Order.module.css';

interface DecodedToken {
    clientid: string;
    role: string;
    exp: number;
}

type SelectedProduct = {
    productid: string;
    name: string;
    price: number;
    quantity: number;
};

type OrderProps = {
    selectedProducts: SelectedProduct[];
    onQuantityChange: (productid: string, quantity: number) => void;
};

const Order: React.FC<OrderProps> = ({ selectedProducts, onQuantityChange }) => {
    const [submitting, setSubmitting] = useState(false);
    const [message, setMessage] = useState('');
    const navigate = useNavigate();
    const location = useLocation();

    const subtotal = selectedProducts.reduce(
        (total, item) => total + item.price * item.quantity,
        0
    );

    const handleSubmit = async () => {
        const token = localStorage.getItem('authToken');
        if (!token) {
            localStorage.setItem('redirectAfterLogin', location.pathname);
            navigate('/login');
            return;
        }
        let clientId = '';
        try {
            const decoded = jwtDecode<DecodedToken>(token);
            clientId = decoded.clientid;
        } catch (error) {
            console.error("Erro ao decodificar o token:", error);
            navigate('/login');
            return;
        }

        setSubmitting(true);
        setMessage('');

        const orderPayload = {
            clientid: clientId,
            products: selectedProducts.map((item) => ({
                productid: item.productid,
                quantity: item.quantity,
            })),
        };

        try {
            const response = await fetch('http://localhost:8081/orders/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify(orderPayload),
            });

            if (response.ok) {
                setMessage('Pedido enviado com sucesso!');
            } else {
                setMessage('Erro ao enviar pedido.');
            }
        } catch (error) {
            setMessage('Erro na comunicação com a API.');
            console.error(error);
        }

        setSubmitting(false);
    };

    return (
        <div className={styles.container}>
            <h3 className={styles.title}>Your Orders</h3>
            {selectedProducts.map((item) => (
                <div key={item.productid} className={styles.item}>
                    <input
                        type="number"
                        min={0}
                        value={item.quantity}
                        onChange={(e) =>
                            onQuantityChange(item.productid, parseInt(e.target.value, 10) || 0)
                        }
                        className={styles.input}
                    />
                    <div className={styles.details}>
                        <div>{item.name}</div>
                        <div className={styles.price}>R${item.price}</div>
                    </div>
                </div>
            ))}
            <hr />
            <div className={styles.subtotal}>
                <span>Subtotal</span>
                <span>R${subtotal}</span>
            </div>
            <button
                className={styles.button}
                disabled={submitting}
                onClick={handleSubmit}
            >
                {submitting ? 'ENVIANDO...' : `CONTINUAR (R$${subtotal})`}
            </button>
            {message && <p className={styles.message}>{message}</p>}
        </div>
    );
};

export default Order;
