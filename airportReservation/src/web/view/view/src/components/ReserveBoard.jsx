import Seat from "./Seat";
import "./styles.css";

export default function ReserveBoard({ seats, onSeatSelect }) {
  return (
    <div className="reserve-board">
      {seats.map((seat, index) => (
        <Seat
          key={index}
          seat={seat}
          onSeatSelect={onSeatSelect}
        />
      ))}
    </div>
  );
}
