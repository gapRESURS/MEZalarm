function updateTemperatureAndHumidity() {
    const rows = document.querySelectorAll('tr');

    rows.forEach(row => {
        const temperatureCell = row.querySelector('#temperature');
        const humidityCell = row.querySelector('#humidity');
        const id = row.id;

        const temperatureUrl = `/API/${id}/DHT22`;
        const humidityUrl = `/API/${id}/DHT22`;

        Promise.all([
            fetch(temperatureUrl),
            fetch(humidityUrl)
        ])
            .then(responses => Promise.all(responses.map(response => response.json())))
            .then(data => {
                if (data[0] && data[1]) {
                    temperatureCell.textContent = `${data[0].temperature}°C`;
                    humidityCell.textContent = `${data[1].humidity}%`;
                } else {
                    console.error('Ошибка при получении данных о температуре или влажности');
                }
            })
            .catch(error => console.error('Ошибка при запросе:', error));
    });
}
updateTemperatureAndHumidity();
setInterval(updateTemperatureAndHumidity, 10000);
window.onload = updateTemperatureAndHumidity;
