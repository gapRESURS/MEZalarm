const ipElements = Array.from(document.querySelectorAll('div[id="ping"]'));

// Функция для выполнения пинга
async function pingIP(element) {
    const ip = element.getAttribute('data-ip');

    try {
        const response = await fetch(`/ping/?ip=${ip}`);
        const data = await response.json();

        if (data.result === true) {
            // IP адрес доступен, меняем цвет текста на зеленый (Bootstrap: success)
            element.classList.remove('text-danger', 'text-secondary'); // Удаляем предыдущие классы
            element.classList.add('text-success');
        } else {
            // IP адрес недоступен, меняем цвет текста на красный (Bootstrap: danger)
            element.classList.remove('text-success', 'text-secondary'); // Удаляем предыдущие классы
            element.classList.add('text-danger');
        }
    } catch (error) {
        console.error('Ошибка при выполнении пинга:', error);
        // Если запрос не удался, меняем цвет текста на серый (Bootstrap: secondary)
        element.classList.remove('text-success', 'text-danger'); // Удаляем предыдущие классы
        element.classList.add('text-secondary');
    }
}

setInterval(() => {
    ipElements.forEach(element => pingIP(element));
}, 1000);