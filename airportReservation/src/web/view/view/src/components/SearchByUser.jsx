import { useState, useEffect } from 'react';
import { getByUser } from '../util/utils';
import SeatIcon from './icons/seat';


export default function SearchByUser() {
    const [userId, setUserId] = useState("");
    const [user, setUser] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(false);

    const handleSearch = async () => {
        setLoading(true);
        setError(null);
        setUser(null);

        try {
            const data = getByUser(userId);
            data.then((res) => {
                if (res.data.length > 0) setUser(res.data)
                else
                    setError("Usuario no encontrado");

            });
        } catch (error) {
            setError("Usuario no encontrado");
        }

        setLoading(false);
    }

    return (
        <div className="search-by-user">
            <div className="search-by-user__input">
                <label htmlFor="user-id">Ingrese su ID de usuario</label>
                <input type="text" placeholder="Ingrese Id de usuario" className="modal-body__seat__input" value={userId} onChange={(e) => setUserId(e.target.value)} />
            </div>
            <div className="search-by-user__button">
                <button className="modal-actions__seat__button" onClick={handleSearch}>Buscar</button>
            </div>
            {loading && <div>Cargando...</div>}
            {error && <div>{error}</div>}
            {user && <div>Usuario encontrado: {userId}</div>}
            {
                user && (
                    <div className="search-by-user__seats">
                        {
                            user.map((seat) => (
                                <div className='seat-row'>
                                    <SeatIcon state={seat.payment} />
                                    <span>Asiento: {seat.chair.chair_id}</span>
                                    <span>Estado: {seat.payment}</span>
                                </div>
                            ))
                        }
                    </div>
                )
            }
        </div>
    );
}