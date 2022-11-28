import SeatIcon from "./icons/seat";
import { useState } from "react";
import { doReserve, doPayment, doDelete } from "../util/utils";

function ModalBody({ isReserved, chair_id, ticket, closeModal, user }) {
    const [ userId, setUserId ] = useState("");

    const message = isReserved === "pending" ? "realizar pago de la reserva?" : "reservar el asiento?";
    const action = isReserved === "pending" ? "Pagar" : "Reservar";

    const handleAction = () => {
        if (isReserved === "available") {
            doReserve( parseInt(userId), chair_id );
            closeModal();
            return;
        }

        doPayment( ticket, chair_id );
        closeModal();

    }

    const deleteAction = () => {
        doDelete( ticket, user );
        closeModal();
    }

    return (
        <div className="modal-body">
            <div className="modal-body__seat">
                <span className="modal-body__seat__number">
                   ¿Desea  { message }
                </span>
                { isReserved !== "pending" && (
                    <div className="user-id-input">
                        <label htmlFor="user-id">Ingrese su ID de usuario</label>
                        <input type="text" placeholder="Ingrese Id de usuario" className="modal-body__seat__input" value={userId} onChange={(e) => setUserId(e.target.value)} />
                    </div>
                ) }
            </div>

            <div className="modal-actions__seat">
                { isReserved === "pending" && <button className="modal-actions__seat__button" onClick={deleteAction}> Cancelar Reserva </button> }
                <button className="modal-actions__seat__button" onClick={handleAction}> { action } </button>    
            </div>
        </div>
    );
}

export default function Modal({ selected, closeModal }) {

    const { isReserved, chair_id, ticket, user } = selected;

    return (
        <div className="modal">
            <div className="modal-bg" ></div>
            <div className="modal-content">
                <div className="modal-header">
                    <span className="modal-header__title">
                        <SeatIcon state={isReserved} />
                        Seat {chair_id}
                    </span>
                    <button className="modal-header__close" onClick={closeModal}>
                        Cerrar
                    </button>
                </div>
                <div className="modal-body">
                    <ModalBody {...{ isReserved, chair_id, ticket, closeModal, user }} />
                </div>
            </div>
        </div>
    );
}