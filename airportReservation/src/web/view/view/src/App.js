import logo from './logo.svg';
import './App.css';
import { useEffect, useState } from 'react';
import axios from 'axios';
import { proccesData } from './util/utils';
import ReserveBoard from './components/ReserveBoard';
import Modal from './components/Modal';
import SearchByUser from './components/SearchByUser';

function App() {

  const [data, setData] = useState([]);
  const [selected, setSelected] = useState(null);
  const [fetching, setFetching] = useState(false);

  useEffect(() => {
    const fetchData = () => {
      axios.get("http://localhost:8080/api/v1/list")
        .then(res => {
          setData(proccesData(res.data));
          setFetching(false);
        })
        .catch(err => {
          console.log(err);
        });
    }

    fetchData();

  }, [fetching, selected]);

  const onSelectSeat = (seat) => {
    if(seat.isReserved === "Accepted") return;
    setSelected(data.find(s => s.chair_id === seat.chair_id));
  }

  return (
    <div className="App">
      <h2>System Reservation</h2>
      <ReserveBoard seats={data} onSeatSelect={onSelectSeat} />
      { selected && <Modal selected={selected} closeModal={() => setSelected("")} /> }

      <SearchByUser />
    </div>
  );
}

export default App;
