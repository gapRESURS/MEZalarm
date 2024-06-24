function updateRelayStatus() {
    const rows = document.querySelectorAll('tr');

    rows.forEach(row => {
        const relayCell = row.querySelector('#relay');
        const id = row.id;

        const url = `/API/${id}/RelayStatus`;

        fetch(url)
            .then(response => response.json())
            .then(data => {
                if (data.result) {
                    relayCell.textContent = 'ВКЛЮЧЕННО';
                } else {
                    relayCell.textContent = 'Выключено';
                }
            })
            .catch(error => console.error('Ошибка при запросе:', error));
    });
}

updateRelayStatus();
setInterval(updateRelayStatus, 1000);
window.onload = updateRelayStatus;