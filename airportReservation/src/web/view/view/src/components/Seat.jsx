import "./styles.css";
import SeatIcon from "./icons/seat";

export default function Seat({ seat, onSeatSelect }) {
  const seatClass =  "seat " + seat.isReserved;

  return (
    <div
      className={seatClass}
      onClick={() => onSeatSelect(seat)}
    >
      <SeatIcon state={seat.isReserved} />
    </div>
  );
}
