// Объединенная функция для обновления состояния реле и температуры/влажности
function updateSystemState() {
    // Вызов функции для обновления состояния реле
    updateRelayStatus();

    // Вызов функции для обновления температуры и влажности
    updateTemperatureAndHumidity();
}

// Функция для обновления статуса реле
function updateRelayStatus() {
    const rows = document.querySelectorAll('tr');

    rows.forEach(row => {
        const relayCell = row.querySelector('#relay');
        const id = row.id;

        const url = `/API/${id}/RelayStatus`;

        fetch(url)
            .then(response => {
                if (!response.ok || response.status === 0) {
                    throw new Error('Ошибка при получении данных');
                }
                return response.json();
            })
            .then(response => {
                if (response.result) {
                    relayCell.textContent = 'ВКЛЮЧЕННО';
                } else {
                    relayCell.textContent = 'Выключено';
                }
            })
            .catch(error => {
                console.error('Ошибка при запросе:', error);
                relayCell.textContent = '-';
            });
    });
}

// Функция для обновления температуры и влажности
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
            .then(async responses => {
                try {
                    const [temperatureResponse, humidityResponse] = await Promise.all(responses.map(response =>
                        response.json().catch(() => null)
                    ));

                    // Проверяем наличие ошибок в ответах
                    if (!temperatureResponse ||!humidityResponse) {
                        throw new Error("Ошибка при получении данных");
                    }

                    // Устанавливаем значения ячеек
                    temperatureCell.textContent = temperatureResponse.temperature? `${temperatureResponse.temperature}°C` : "-";
                    humidityCell.textContent = humidityResponse.humidity? `${humidityResponse.humidity}%` : "-";
                } catch (error) {
                    console.error('Ошибка при запросе:', error);
                    temperatureCell.textContent = "-";
                    humidityCell.textContent = "-";
                }
            })
            .catch(error => {
                console.error('Общая ошибка при обновлении температуры и влажности:', error);
                temperatureCell.textContent = "-";
                humidityCell.textContent = "-";
            });
    });
}

// Запускаем функцию updateSystemState сразу после загрузки страницы и каждую секунду
updateSystemState();
setInterval(updateSystemState, 3000);
