import axios from "axios";

export const proccesData = (data) => {
    const result = orderData(data);
    const toShow = JSON.parse(sessionStorage.getItem("seats")) || result;

    result.forEach((item) => {

        toShow.forEach((seat) => {
            item.isReserved = "available";
            if (item.chair_id === seat.chair_id) {
                seat.isReserved = seat.isReserved || item.isReserved;
            }
        });
    });

    sessionStorage.setItem("seats", JSON.stringify(toShow));

    return toShow;
}

export const orderData = (data) => {
    const result = data.data;

    result.sort((a, b) => {
        return a.price - b.price;
    });

    return result;
}


export const doReserve = async (user_id, chair_id) => {
    fetch("http://localhost:8080/api/v1/reserve", {
        method: "POST",
        body: JSON.stringify({
            user_id,
            chair_id
        })
    })
    .then(res => res.json())
    .then(res => {
        const seats = JSON.parse(sessionStorage.getItem("seats"));
        const index = seats.findIndex(s => s.chair_id === chair_id);

        seats[index].isReserved = "pending";
        const { data } = res;

        seats[index].ticket = data.ticket_id;
        seats[index].user = data.user_id;
        sessionStorage.setItem("seats", JSON.stringify(seats));
        console.log(seats)
    })
}

export const doPayment = async ( ticket_id, chair_id ) => {
    fetch(`http://localhost:8080/api/v1/pay/${ticket_id}`, {
        method: "POST"
    })
    .then(res => {
        const seats = JSON.parse(sessionStorage.getItem("seats"));
        const index = seats.findIndex(s => s.chair_id === chair_id);

        seats[index].isReserved = "Accepted";
        sessionStorage.setItem("seats", JSON.stringify(seats));
        console.log(seats)
    })
}

export const getByUser = async (user_id) => {
    const res = await axios.get(`http://localhost:8080/api/v1/list/${user_id}`);
    return res.data;
}

export const doDelete = async (ticket_id, user_id) => {
    fetch(`http://localhost:8080/api/v1/delete`, {
        method: "DELETE",
        body: JSON.stringify({
            ticket_id,
            user_id
        })
    })
    .then(res => {
        const seats = JSON.parse(sessionStorage.getItem("seats"));
        const index = seats.findIndex(s => s.ticket === ticket_id);

        seats[index].isReserved = "available";
        seats[index].user = null;
        sessionStorage.setItem("seats", JSON.stringify(seats));
        console.log(seats)
    })
}